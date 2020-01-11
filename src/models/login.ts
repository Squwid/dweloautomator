import { User } from "../models/user";

interface LoginRequest {
  email: string;
  password: string;
  applicationId: string;
}

interface LoginResponse {
  token: string;
  user: User;
}
