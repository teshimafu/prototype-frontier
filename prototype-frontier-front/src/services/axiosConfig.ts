import axios, { AxiosInstance } from "axios";
import FirebaseService from "./firebaseService";

export const axiosInstance: AxiosInstance = axios.create(
  process.env.NODE_ENV === "development"
    ? { baseURL: "http://localhost:3000" }
    : {}
);

export const axiosAuthInstance: Promise<AxiosInstance> = new Promise(
  resolve => {
    FirebaseService.getInstance()
      .getIdToken()
      .then(token =>
        resolve(
          axios.create({
            baseURL:
              process.env.NODE_ENV === "development"
                ? "http://localhost:3000"
                : "",
            headers: {
              Authorization: `Bearer ${token ?? ""}`
            }
          })
        )
      );
  }
);
