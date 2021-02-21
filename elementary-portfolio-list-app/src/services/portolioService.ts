import axios from "axios";
import { Portfolio } from "@/models/Portfolio";

export class PortfolioService {
  /**
   * getPortfolio
   */
  public static getPortfolio = (): Portfolio | null => {
    return null;
  };

  /**
   * getPortfolio
   */
  public static getPortfolioList = async (): Promise<Portfolio[]> => {
    //let { data } = await axios.get<Portfolio[]>("/api/portfolio");
    // APIリクエストのモック 2000ms後にユーザー名の配列を返す
    const fetchUsers = () => {
      return new Promise<Portfolio[]>((resolve) => {
        setTimeout(() => {
          resolve([
            {
              id: 1,
              title: "シロナガスクジラ",
              author: "teshima",
              createDate: "1/1",
              tags: ["1", "2", "3"],
              abstruct: "## title etc...",
              source: "https:google.com",
              url: "https://qiita.com"
            }
          ]);
        }, 2000);
      });
    }; // ユーザー名を取得

    // setup()内で読んでいるのでVue.js 2系で言うonCreated()と同じライフサイクルで実施される。
    return fetchUsers();
  };
}
