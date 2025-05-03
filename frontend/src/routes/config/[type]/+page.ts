import type {PageLoad} from './$types';

export type Config = {
    type: string;
    config: string;
}

export const load: PageLoad = ({params, fetch}) => {
    return {
        loader: async function() {
            let response = await fetch(`/api/config/${params.type}`);
            let text = await response.text();
            return <Config>JSON.parse(text);
        }
    }
};