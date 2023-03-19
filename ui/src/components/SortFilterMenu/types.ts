import { SearchSortMethod } from "@/store/search/types";

export type FormValues = {
  sortingMethod: SearchSortMethod;
  typologyNames: string[];
  subTypologyNames: string[];
};
