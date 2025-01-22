import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { format } from "date-fns";
import { Video } from "../types";

type VideoCardProps = {
  video: Video;
};

export const VideoCard = ({ video }: VideoCardProps) => {
  return (
    <Card className="w-full">
      <CardHeader className="space-y-2">
        <div className="aspect-video w-full overflow-hidden rounded-lg">
          <img
            src={video.thumbnailUrl}
            alt={video.title}
            className="w-full h-full object-cover"
          />
        </div>
        <CardTitle className="text-xl">{video.title}</CardTitle>
        <CardDescription>
          {video.channel} • {format(new Date(video.publishedAt), "PP")}
        </CardDescription>
      </CardHeader>
      <CardContent>
        <p className="text-sm text-muted-foreground line-clamp-2">
          {video.description}
        </p>
      </CardContent>
    </Card>
  );
};
