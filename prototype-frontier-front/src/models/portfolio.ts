import { PartiallyPartial } from "./customTSlib";

export interface Portfolio {
  id: number;
  title: string;
  author: string;
  created_at: string;
  abstruct: string;
  link: string;
  source: string;
  readme: string;
}
export type InputPortfolio = PartiallyPartial<Portfolio, "id" | "created_at">;
