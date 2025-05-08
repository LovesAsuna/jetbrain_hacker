import type {PageLoad} from './$types';

export interface ProductDto {
    code: string,
    salesCode: string,
    name: string,
    description: string,
    forSale: boolean
    productFamilyName: string
    releases: Array<{ version: string }>
}

export interface Product {
    code: string,
    name: string,
    productFamilyName: string
    version: string
}

const DataBaseUrl = "https://data.services.jetbrains.com"
export const load: PageLoad = async ({fetch}) => {
    let resp = await fetch(`${DataBaseUrl}/products?fields=name,code,forSale,salesCode,description,productFamilyName,releases.version`)
    let productDtos: Array<ProductDto> = await resp.json()
    let products: Array<Product> = [];
    productDtos.forEach(dto => {
        dto.productFamilyName = dto.productFamilyName.replace(" ", "-").toLowerCase()
        let extra = dto.forSale && (dto.salesCode != dto.code);
        products.push(
            {
                code: dto.code,
                name: extra ? `${dto.name}(${dto.code})` : dto.name,
                productFamilyName: dto.productFamilyName,
                version: dto.releases.length > 0 ? dto.releases[0].version: ""
            }
        );
        if (extra) {
            products.push(
                {
                    code: dto.salesCode,
                    name: `${dto.name}(${dto.salesCode})`,
                    productFamilyName: dto.productFamilyName,
                    version: dto.releases.length > 0 ? dto.releases[0].version: ""
                }
            );
        }
    });
    return {
        products
    }
};