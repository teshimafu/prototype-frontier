<script lang="ts">
import { Portfolio } from "../models/portfolio";
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
   *postPortfolioList
   */
  public static postPortfolio = async (portfolio: Portfolio) => {
    await axiosInstance.post("/api/portfolio", portfolio);
  };
}
</script>
