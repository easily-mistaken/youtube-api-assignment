export type Video = {
  id: number;
  title: string;
  description: string;
  publishedAt: string;
  thumbnailUrl: string;
  channel: string;
};

export type ApiResponse = {
  limit: number;
  page: number;
  videos: Video[];
};
