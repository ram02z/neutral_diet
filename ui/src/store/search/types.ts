export enum SearchType {
  History,
  Global,
}

export enum SearchSortMethod {
  NameDescending,
  NameAscending,
  EmissionsDescending,
  EmissionsAscending,
}

export type SearchFilters = {
  typologies: string[];
  subTypologies: string[];
};
