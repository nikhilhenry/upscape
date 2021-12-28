import axios from "axios";
import {
  Form,
  Response,
  User,
  UserRequest,
} from "../types/userTypes.interface";

export const login = async (form: Form): Promise<Response> => {
  try {
    const response: Response = await (
      await axios.post("/api/auth/login", form)
    ).data;

    return response;
  } catch (error) {
    const response: Response = {
      success: false,
      message: "email or password is incorrect",
    };

    return response;
  }
};

export const getMeta = async (): Promise<User> => {
  try {
    const userMeta: User = await (await axios.get("/api/user")).data;
    return userMeta;
  } catch (error) {
    console.log(error);
    return { name: "", img_url: "" };
  }
};

export const updateMeta = async (
  payload: UserRequest
): Promise<User | null> => {
  try {
    const user: User = await (await axios.put("/api/user", payload)).data;
    return user;
  } catch (error) {
    console.log(error);
    return null;
  }
};
