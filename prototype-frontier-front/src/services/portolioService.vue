<script lang="ts">
import { Portfolio, InputPortfolio } from "../models/portfolio";
import { axiosInstance } from "../services/axiosConfig";

export default class PortfolioService {
  /**
   * getPortfolio
   */
  public static getPortfolio = async (id: number): Promise<Portfolio> => {
    const { data } = await axiosInstance.get<Portfolio>("/api/portfolio/" + id);
    return data;
  };

  /**
   * getPortfolioList
   */
  public static getPortfolioList = async (): Promise<Portfolio[]> => {
    const { data } = await axiosInstance.get<Portfolio[]>("/api/portfolio");
    return data;
  };

  /**
   *postPortfolio
   */
  public static postPortfolio = async (portfolio: InputPortfolio) => {
    return await axiosInstance.post<Portfolio>("/api/portfolio", portfolio);
  };

  /**
   *putPortfolio
   */
  public static putPortfolio = async (portfolio: Portfolio) => {
    return await axiosInstance.put<Portfolio>("/api/portfolio", portfolio);
  };
}
</script>
