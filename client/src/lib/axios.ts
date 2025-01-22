import axios from "axios";
import { ApiResponse } from "../types";

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
});

export const sleep = (ms: number) =>
  new Promise((resolve) => setTimeout(resolve, ms));
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
  //   return response.data;

  await sleep(3000);
  return {
    limit: 9,
    page: 1,
    videos: [
      {
        id: 4801,
        title:
          "JERSEY CONTROVERSY 😱 PAK NAME LIKHNA PADEGA KYA 😎 TEAM INDIA\u0026#39;S PLAYING 11 🔥 SKYBALL Vs BAZZBALL",
        description:
          "JERSEY CONTROVERSY PAK NAME LIKHNA PADEGA KYA TEAM INDIA'S PLAYING 11 SKYBALL Vs BAZZBALL ...",
        publishedAt: "2025-01-22T06:07:57+05:30",
        thumbnailUrl: "https://i.ytimg.com/vi/FVul6P_aHME/hqdefault_live.jpg",
        channel: "AB Cricinfo",
      },
      {
        id: 4751,
        title:
          "JERSEY CONTROVERSY 😱 PAK NAME LIKHNA PADEGA KYA 😎 TEAM INDIA\u0026#39;S PLAYING 11 🔥 SKYBALL Vs BAZZBALL",
        description:
          "JERSEY CONTROVERSY PAK NAME LIKHNA PADEGA KYA TEAM INDIA'S PLAYING 11 SKYBALL Vs BAZZBALL ...",
        publishedAt: "2025-01-22T06:07:57+05:30",
        thumbnailUrl: "https://i.ytimg.com/vi/FVul6P_aHME/hqdefault_live.jpg",
        channel: "AB Cricinfo",
      },
      {
        id: 4701,
        title:
          "JERSEY CONTROVERSY 😱 PAK NAME LIKHNA PADEGA KYA 😎 TEAM INDIA\u0026#39;S PLAYING 11 🔥 SKYBALL Vs BAZZBALL",
        description:
          "JERSEY CONTROVERSY PAK NAME LIKHNA PADEGA KYA TEAM INDIA'S PLAYING 11 SKYBALL Vs BAZZBALL ...",
        publishedAt: "2025-01-22T06:07:57+05:30",
        thumbnailUrl: "https://i.ytimg.com/vi/FVul6P_aHME/hqdefault_live.jpg",
        channel: "AB Cricinfo",
      },
      {
        id: 4651,
        title:
          "JERSEY CONTROVERSY 😱 PAK NAME LIKHNA PADEGA KYA 😎 TEAM INDIA\u0026#39;S PLAYING 11 🔥 SKYBALL Vs BAZZBALL",
        description:
          "JERSEY CONTROVERSY PAK NAME LIKHNA PADEGA KYA TEAM INDIA'S PLAYING 11 SKYBALL Vs BAZZBALL ...",
        publishedAt: "2025-01-22T06:07:57+05:30",
        thumbnailUrl: "https://i.ytimg.com/vi/FVul6P_aHME/hqdefault_live.jpg",
        channel: "AB Cricinfo",
      },
      {
        id: 4601,
        title:
          "JERSEY CONTROVERSY 😱 PAK NAME LIKHNA PADEGA KYA 😎 TEAM INDIA\u0026#39;S PLAYING 11 🔥 SKYBALL Vs BAZZBALL",
        description:
          "JERSEY CONTROVERSY PAK NAME LIKHNA PADEGA KYA TEAM INDIA'S PLAYING 11 SKYBALL Vs BAZZBALL ...",
        publishedAt: "2025-01-22T06:07:57+05:30",
        thumbnailUrl: "https://i.ytimg.com/vi/FVul6P_aHME/hqdefault_live.jpg",
        channel: "AB Cricinfo",
      },
      {
        id: 4551,
        title:
          "JERSEY CONTROVERSY 😱 PAK NAME LIKHNA PADEGA KYA 😎 TEAM INDIA\u0026#39;S PLAYING 11 🔥 SKYBALL Vs BAZZBALL",
        description:
          "JERSEY CONTROVERSY PAK NAME LIKHNA PADEGA KYA TEAM INDIA'S PLAYING 11 SKYBALL Vs BAZZBALL ...",
        publishedAt: "2025-01-22T06:07:57+05:30",
        thumbnailUrl: "https://i.ytimg.com/vi/FVul6P_aHME/hqdefault_live.jpg",
        channel: "AB Cricinfo",
      },
      {
        id: 4501,
        title:
          "JERSEY CONTROVERSY 😱 PAK NAME LIKHNA PADEGA KYA 😎 TEAM INDIA\u0026#39;S PLAYING 11 🔥 SKYBALL Vs BAZZBALL",
        description:
          "JERSEY CONTROVERSY PAK NAME LIKHNA PADEGA KYA TEAM INDIA'S PLAYING 11 SKYBALL Vs BAZZBALL ...",
        publishedAt: "2025-01-22T06:07:57+05:30",
        thumbnailUrl: "https://i.ytimg.com/vi/FVul6P_aHME/hqdefault_live.jpg",
        channel: "AB Cricinfo",
      },
      {
        id: 4451,
        title:
          "JERSEY CONTROVERSY 😱 PAK NAME LIKHNA PADEGA KYA 😎 TEAM INDIA\u0026#39;S PLAYING 11 🔥 SKYBALL Vs BAZZBALL",
        description:
          "JERSEY CONTROVERSY PAK NAME LIKHNA PADEGA KYA TEAM INDIA'S PLAYING 11 SKYBALL Vs BAZZBALL ...",
        publishedAt: "2025-01-22T06:07:57+05:30",
        thumbnailUrl: "https://i.ytimg.com/vi/FVul6P_aHME/hqdefault_live.jpg",
        channel: "AB Cricinfo",
      },
      {
        id: 4401,
        title:
          "JERSEY CONTROVERSY 😱 PAK NAME LIKHNA PADEGA KYA 😎 TEAM INDIA\u0026#39;S PLAYING 11 🔥 SKYBALL Vs BAZZBALL",
        description:
          "JERSEY CONTROVERSY PAK NAME LIKHNA PADEGA KYA TEAM INDIA'S PLAYING 11 SKYBALL Vs BAZZBALL ...",
        publishedAt: "2025-01-22T06:07:57+05:30",
        thumbnailUrl: "https://i.ytimg.com/vi/FVul6P_aHME/hqdefault_live.jpg",
        channel: "AB Cricinfo",
      },
    ],
  };
};
