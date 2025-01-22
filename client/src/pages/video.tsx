import { useQuery } from "@tanstack/react-query";
import { useSearchStore } from "../store";
import { fetchVideos } from "../lib/axios";
import { Input } from "../components/ui/input";
import { VideoCard } from "../components/video-card";
import { useState } from "react";
import { Loader2 } from "lucide-react";

export const VideoPage = () => {
  const [page, setPage] = useState(1);
  const { searchQuery, setSearchQuery } = useSearchStore();

  const { data, isLoading, isError } = useQuery({
    queryKey: ["videos", page, searchQuery],
    queryFn: () => fetchVideos(page, 10, searchQuery),
  });

  const handleSearch = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchQuery(e.target.value);
    setPage(1); // Reset to first page when searching
  };

  if (isError) {
    return (
      <div className="flex items-center justify-center h-screen">
        <p className="text-red-500">
          Error loading videos. Please try again later.
        </p>
      </div>
    );
  }

  return (
    <div className="container mx-auto py-8 space-y-6">
      <div className="max-w-xl mx-auto">
        <Input
          type="search"
          placeholder="Search videos..."
          value={searchQuery}
          onChange={handleSearch}
          className="w-full"
        />
      </div>

      {isLoading ? (
        <div className="flex items-center justify-center py-12">
          <Loader2 className="h-8 w-8 animate-spin" />
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {data?.videos?.length !== 0 ? (
            data?.videos.map((video) => (
              <VideoCard key={video.id} video={video} />
            ))
          ) : (
            <p>No videos found</p>
          )}
        </div>
      )}
    </div>
  );
};
