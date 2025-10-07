export interface RegisterRequest {
  name: string;
  surname: string;
  username: string;
  password: string;
}

export interface RegisterResponse {
  id: string;
  name: string;
  surname: string;
  username: string;
}

export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  token: string;
  user: {
    id: string;
    name: string;
    username: string;
  };
}
