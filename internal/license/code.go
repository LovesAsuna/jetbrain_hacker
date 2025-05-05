package license

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lovesasuna/sync/coroutinegroup"
	"net/http"
)

const (
	DataBaseUrl   = "https://data.services.jetbrains.com"
	PluginBaseUrl = "https://plugins.jetbrains.com"
)

type ProductDto struct {
	Code        string `json:"code"`
	SalesCode   string `json:"salesCode"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ForSale     bool   `json:"forSale"`
}

func GetProductCode() (codes []string, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/products?fields=name,code,forSale,salesCode,description", DataBaseUrl))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var productList []ProductDto
	if err = json.NewDecoder(resp.Body).Decode(&productList); err != nil {
		return nil, err
	}
	for _, product := range productList {
		codes = append(codes, product.Code)
		if product.ForSale && product.SalesCode != "" {
			codes = append(codes, product.SalesCode)
		}
	}
	return
}

type PluginDto struct {
	ID           int32    `json:"id"`
	XMLID        string   `json:"xmlId"`
	Link         string   `json:"link"`
	Name         string   `json:"name"`
	Preview      string   `json:"preview"`
	Downloads    int      `json:"downloads"`
	PricingModel string   `json:"pricingModel"`
	Icon         string   `json:"icon"`
	PreviewImage string   `json:"previewImage"`
	Cdate        int64    `json:"cdate"`
	Rating       float64  `json:"rating"`
	HasSource    bool     `json:"hasSource"`
	Tags         []string `json:"tags"`
	Vendor       Vendor   `json:"vendor"`
}

type PluginDetail struct {
	ID                      int32        `json:"id"`
	Name                    string       `json:"name"`
	Link                    string       `json:"link"`
	Approve                 bool         `json:"approve"`
	XMLID                   string       `json:"xmlId"`
	Description             string       `json:"description"`
	CustomIdeList           bool         `json:"customIdeList"`
	Preview                 string       `json:"preview"`
	DocText                 string       `json:"docText"`
	Cdate                   int64        `json:"cdate"`
	Family                  string       `json:"family"`
	Downloads               int          `json:"downloads"`
	PurchaseInfo            PurchaseInfo `json:"purchaseInfo"`
	Vendor                  Vendor       `json:"vendor"`
	Urls                    Urls         `json:"urls"`
	Tags                    []Tags       `json:"tags"`
	HasUnapprovedUpdate     bool         `json:"hasUnapprovedUpdate"`
	PricingModel            string       `json:"pricingModel"`
	Screens                 []Screens    `json:"screens"`
	Icon                    string       `json:"icon"`
	IsHidden                bool         `json:"isHidden"`
	IsMonetizationAvailable bool         `json:"isMonetizationAvailable"`
	IsBlocked               bool         `json:"isBlocked"`
	IsModificationAllowed   bool         `json:"isModificationAllowed"`
}

type PurchaseInfo struct {
	ProductCode   string      `json:"productCode"`
	BuyURL        interface{} `json:"buyUrl"`
	PurchaseTerms interface{} `json:"purchaseTerms"`
	Optional      bool        `json:"optional"`
	TrialPeriod   int         `json:"trialPeriod"`
}

type Details struct {
	City    string      `json:"city"`
	Address string      `json:"address"`
	State   interface{} `json:"state"`
	Zip     string      `json:"zip"`
	Phone   string      `json:"phone"`
}

type Vendor struct {
	Type        string  `json:"type"`
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	URL         string  `json:"url"`
	Link        string  `json:"link"`
	PublicName  string  `json:"publicName"`
	Email       string  `json:"email"`
	CountryCode string  `json:"countryCode"`
	Country     string  `json:"country"`
	IsVerified  bool    `json:"isVerified"`
	VendorID    int     `json:"vendorId"`
	Details     Details `json:"details"`
	IsTrader    bool    `json:"isTrader"`
}

type Urls struct {
	URL              string `json:"url"`
	ForumURL         string `json:"forumUrl"`
	LicenseURL       string `json:"licenseUrl"`
	PrivacyPolicyURL string `json:"privacyPolicyUrl"`
	BugtrackerURL    string `json:"bugtrackerUrl"`
	DocURL           string `json:"docUrl"`
	SourceCodeURL    string `json:"sourceCodeUrl"`
}

type Tags struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Privileged bool   `json:"privileged"`
	Searchable bool   `json:"searchable"`
	Link       string `json:"link"`
}

type Screens struct {
	URL string `json:"url"`
}

func GetPluginCode(max, offset int32, keyword string) (codes []string, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/searchPlugins?max=%d&offset=%d&search=%s", PluginBaseUrl, max, offset, keyword))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pluginListResp struct {
		Plugins []*PluginDto `json:"plugins"`
	}
	err = json.NewDecoder(resp.Body).Decode(&pluginListResp)
	if err != nil {
		return nil, err
	}

	pluginDetailChan := make(chan *PluginDetail, len(pluginListResp.Plugins))
	group, _ := coroutinegroup.WithContext(context.Background())
	group.SetMaxErrorTask(int32(len(pluginListResp.Plugins) / 3))
	group.SetGlobalRetryTimes(1)
	for _, plugin := range pluginListResp.Plugins {
		if plugin.PricingModel == "FREE" {
			continue
		}
		if plugin.Icon != "" {
			plugin.Icon = PluginBaseUrl + plugin.Icon
		}
		p := plugin
		group.Go(
			func(ctx context.Context) error {
				detail, err := getDetailByPluginId(p.ID)
				if err != nil {
					return err
				}
				pluginDetailChan <- detail
				return nil
			},
		)
	}
	errs := group.Wait()
	close(pluginDetailChan)
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}
	for detail := range pluginDetailChan {
		code := detail.PurchaseInfo.ProductCode
		if code != "" {
			codes = append(codes, code)
		}
	}
	return
}

func getDetailByPluginId(id int32) (detail *PluginDetail, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/plugins/%d", PluginBaseUrl, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	detail = new(PluginDetail)
	err = json.NewDecoder(resp.Body).Decode(detail)
	return
}
