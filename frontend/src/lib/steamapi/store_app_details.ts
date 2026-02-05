
export const GetStoreInfo = async (appID: number, fetch: typeof globalThis.fetch = globalThis.fetch): Promise<AppInfo> => {
    const url = `https://store.steampowered.com/api/appdetails?appids=${appID}`;
    const response = await fetch(url);
    if (!response.ok) {
        const body = await response.text();
        throw new Error(`HTTP error ${response.status}: ${body}`);
    }
    const data: StoreAppDetails = await response.json();
    if (data?.[`${appID}`]?.success) {
        return data[`${appID}`]!.data;
    }

    throw new Error('app details not found in response');
};


export type StoreAppDetails = Record<string, AppDetails>;

export interface AppDetails {
    success: boolean;
    data: AppInfo;
}

export interface AppInfo {
    type: string;
    name: string;
    steam_appid: number;
    required_age: number;
    is_free: boolean;
    controller_support: string;
    dlc: number[];
    detailed_description: string;
    about_the_game: string;
    short_description: string;
    supported_languages: string;
    header_image: string;
    capsule_image: string;
    capsule_imagev5: string;
    website: string;
    pc_requirements: PcRequirements;
    mac_requirements: MacRequirements;
    linux_requirements: unknown[]; // TODO
    developers: string[];
    publishers: string[];
    packages: number[];
    package_groups: PackageGroup[];
    platforms: Platforms;
    metacritic: Metacritic;
    categories: Category[];
    genres: Genre[];
    screenshots: Screenshot[];
    movies: Movie[];
    recommendations: Recommendations;
    achievements: Achievements;
    release_date: ReleaseDate;
    support_info: SupportInfo;
    background: string;
    background_raw: string;
    content_descriptors: ContentDescriptors;
    ratings: Ratings;
}

export interface PcRequirements {
    minimum: string;
}

export interface MacRequirements {
    minimum: string;
}

export interface PackageGroup {
    name: string;
    title: string;
    description: string;
    selection_text: string;
    save_text: string;
    display_type: number;
    is_recurring_subscription: string;
    subs: Sub[];
}

export interface Sub {
    packageid: number;
    percent_savings_text: string;
    percent_savings: number;
    option_text: string;
    option_description: string;
    can_get_free_license: string;
    is_free_license: boolean;
    price_in_cents_with_discount: number;
}

export interface Platforms {
    windows: boolean;
    mac: boolean;
    linux: boolean;
}

export interface Metacritic {
    score: number;
    url: string;
}

export interface Category {
    id: number;
    description: string;
}

export interface Genre {
    id: string;
    description: string;
}

export interface Screenshot {
    id: number;
    path_thumbnail: string;
    path_full: string;
}

export interface Movie {
    id: number;
    name: string;
    thumbnail: string;
    dash_av1: string;
    dash_h264: string;
    hls_h264: string;
    highlight: boolean;
}

export interface Recommendations {
    total: number;
}

export interface Achievements {
    total: number;
    highlighted: Highlighted[];
}

export interface Highlighted {
    name: string;
    path: string;
}

export interface ReleaseDate {
    coming_soon: boolean;
    date: string;
}

export interface SupportInfo {
    url: string;
    email: string;
}

export interface ContentDescriptors {
    ids: unknown[]; // TODO
    notes: string;
}

export interface Ratings {
    usk: Usk;
    dejus: Dejus;
    steam_germany: SteamGermany;
}

export interface Usk {
    rating: string;
}

export interface Dejus {
    rating_generated: string;
    rating: string;
    required_age: string;
    banned: string;
    use_age_gate: string;
    descriptors: string;
}

export interface SteamGermany {
    rating_generated: string;
    rating: string;
    required_age: string;
    banned: string;
    use_age_gate: string;
    descriptors: string;
}
