export interface Form {
  email: string;
  password: string;
}

export interface Response {
  token?: string;
  message: string;
  success: boolean;
}
