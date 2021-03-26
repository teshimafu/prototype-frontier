import { Portfolio, InputPortfolio, SearchQuery } from "../models/portfolio";
import { axiosInstance, axiosAuthInstance } from "../services/axiosConfig";

export default class PortfolioService {
  private static PATH = "/api/portfolios" as const;
  /**
   * getPortfolio
   */
  public static getPortfolio = async (id: number): Promise<Portfolio> => {
    const { data } = await axiosInstance.get<Portfolio>(
      PortfolioService.PATH + "/" + id
    );
    return data;
  };

  /**
   * getPortfolioList
   */
  public static getPortfolioList = async (
    searchQuery?: SearchQuery
  ): Promise<Portfolio[]> => {
    const { data } = await axiosInstance.get<Portfolio[]>(
      PortfolioService.PATH,
      { params: searchQuery }
    );
    return data;
  };

  /**
   * postPortfolio
   */
  public static postPortfolio = async (portfolio: InputPortfolio) => {
    return await (await axiosAuthInstance).post<Portfolio>(
      PortfolioService.PATH,
      portfolio
    );
  };

  /**
   * putPortfolio
   */
  public static putPortfolio = async (id: string, portfolio: Portfolio) => {
    return await (await axiosAuthInstance).put<Portfolio>(
      PortfolioService.PATH + "/" + id,
      portfolio
    );
  };

  /**
   * deletePortfolio
   */
  public static deletePortfolio = async (id: string) => {
    await (await axiosAuthInstance).delete(PortfolioService.PATH + "/" + id);
  };
}
