import { Portfolio } from "@/models/Portfolio";
import { axiosInstance } from "@/services/axiosConfig";

export class PortfolioService {
  /**
   * getPortfolio
   */
  public static getPortfolio = (): Portfolio | null => {
    return null;
  };

  /**
   * getPortfolioList
   */
  public static getPortfolioList = async (): Promise<Portfolio[]> => {
    const { data } = await axiosInstance.get<Portfolio[]>("/api/portfolio");
    return data;
  };
}
