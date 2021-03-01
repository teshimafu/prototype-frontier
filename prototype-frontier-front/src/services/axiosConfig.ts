import axios, { AxiosInstance } from "axios";

export const axiosInstance: AxiosInstance = axios.create(
  process.env.NODE_ENV === "development"
    ? { baseURL: "http://localhost:3000" }
    : {}
);
