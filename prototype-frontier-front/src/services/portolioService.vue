<script lang="ts">
import { Portfolio, InputPortfolio } from "../models/portfolio";
import { axiosInstance } from "../services/axiosConfig";

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
  public static getPortfolioList = async (): Promise<Portfolio[]> => {
    const { data } = await axiosInstance.get<Portfolio[]>(
      PortfolioService.PATH
    );
    return data;
  };

  /**
   * postPortfolio
   */
  public static postPortfolio = async (portfolio: InputPortfolio) => {
    return await axiosInstance.post<Portfolio>(
      PortfolioService.PATH,
      portfolio
    );
  };

  /**
   * putPortfolio
   */
  public static putPortfolio = async (id: string, portfolio: Portfolio) => {
    return await axiosInstance.put<Portfolio>(
      PortfolioService.PATH + "/" + id,
      portfolio
    );
  };

  /**
   * deletePortfolio
   */
  public static deletePortfolio = async (id: string) => {
    await axiosInstance.delete(PortfolioService.PATH + "/" + id);
  };
}
</script>
