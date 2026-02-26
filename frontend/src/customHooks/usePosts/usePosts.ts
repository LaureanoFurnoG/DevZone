import { useEffect, useState } from "react";
import axiosInstance from "../../api/axios";

type TiptapNode = {
  type: string;
  content?: TiptapNode[];
  text?: string;
};

type TiptapDocument = {
  type: "doc";
  content: TiptapNode[];
};

type Category = {
  id: number;
  name: string;
};

type Post = {
  id: number;
  title: string;
  content: TiptapDocument;
  id_user: string;
  username: string;
  profile_image: string;
  created_at: string;
  categoriesdata: Category[];
};

export const usePosts = (Id: number | null) => {
  const [posts, setPosts] = useState<Post[]>([]);
  useEffect(() => {
    const getPosts = async () => {
      try {
        const response = await axiosInstance.get("/devzone-api/v1/posts");
        console.log(response);
        setPosts(response.data.posts);
      } catch (error) {
        console.log(error);
      }
    };

    const getPostsByID = async (Id: number) => {
      try {
        const response = await axiosInstance.get(`/devzone-api/v1/posts/${Id}`);
        console.log(response);
        setPosts(response.data.posts);
      } catch (error) {
        console.log(error);
      }
    };

    if (Id !== null) {
      getPostsByID(Id);
    } else {
      getPosts();
    }
  }, [Id]);
  return { posts, setPosts };
};
