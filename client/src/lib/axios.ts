import axios from "axios";
import { ApiResponse } from "../types";

const api = axios.create({
  baseURL: process.env.VITE_API_URL,
});

export const fetchVideos = async (
  page: number,
  limit: number,
  search?: string
) => {
  const response = await api.get<ApiResponse>("/videos", {
    params: {
      page,
      limit,
      search,
    },
  });
  return response.data;
};
