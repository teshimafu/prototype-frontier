import axios, { AxiosInstance } from "axios";
import restAPIConfig from "../../restAPIConfig.json";
import FirebaseService from "./firebaseService";

export const axiosInstance: AxiosInstance = axios.create({
  baseURL: restAPIConfig.baseURL
});

export const axiosAuthInstance: Promise<AxiosInstance> = new Promise<
  AxiosInstance
>(resolve => {
  FirebaseService.getInstance()
    .getIdToken()
    .then(token =>
      resolve(
        axios.create({
          baseURL: restAPIConfig.baseURL,
          headers: {
            Authorization: `Bearer ${token ?? ""}`
          }
        })
      )
    );
}).then(instance => {
  instance.interceptors.response.use(
    r => r,
    error => {
      if (error.response.status === 401) {
        alert("認証エラーです\n画面更新してやり直してください");
      }
    }
  );
  return instance;
});
