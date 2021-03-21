import axios, { AxiosInstance } from "axios";
import restAPIConfig from "../../restAPIConfig.json";
import FirebaseService from "./firebaseService";

export const axiosInstance: AxiosInstance = axios.create({
  baseURL: restAPIConfig.baseURL
});

export const axiosAuthInstance: Promise<AxiosInstance> = new Promise(
  resolve => {
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
  }
);
