import axios from "axios";
import { Form, Response } from "../types/userTypes.interface";

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
