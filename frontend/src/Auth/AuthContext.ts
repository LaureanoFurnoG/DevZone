import { createContext } from "react";

export type Me = {
  name: string;
  lastname: string;
  email?: string;
  nickName?: string;
  profileImage?: string;
  sub: string | undefined;
  token: string | undefined;
};

export interface AuthContextType {
  isAuthenticated: boolean;
  isLoading: boolean;
  me: Me | null;
  token?: string;
  login: () => void;
  logout: () => void;
}

export const AuthContext = createContext<AuthContextType | null>(null);
