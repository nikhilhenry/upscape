import axios from "axios";
import { Form, Response, User } from "../types/userTypes.interface";

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
    return { name: "", imgUrl: "" };
  }
};
