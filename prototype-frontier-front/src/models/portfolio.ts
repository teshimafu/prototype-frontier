import { PartiallyPartial } from "./customTSlib";

export interface Portfolio {
  id: number;
  title: string;
  uid: string;
  author: string;
  created_at: string;
  abstruct: string;
  link: string;
  source: string;
  readme: string;
}

export interface SearchQuery {
  title: string;
  author: string;
  readme: string;
}

export type InputPortfolio = PartiallyPartial<
  Portfolio,
  "id" | "uid" | "created_at"
>;
