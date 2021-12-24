export interface Form {
  password: string;
}

export interface Response {
  token?: string;
  message: string;
  success: boolean;
}

export interface User {
  name: string;
  img_url: string;
}
