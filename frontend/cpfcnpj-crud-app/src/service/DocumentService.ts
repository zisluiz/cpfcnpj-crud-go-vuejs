import Query from "@/model/Query";
import Response from "@/model/PaginatedResponse";
import Messages from "@/model/Messages";
import Document from "@/model/Document";
import axios from "axios";

export default class DocumentService {
  async findDocumentsBy(query: Query): Promise<Response | Messages> {
    return axios
      .get(`${process.env.VUE_APP_API_HOST}/api/documents`, {
        params: query.toParam(),
      })
      .then((res) => res.data)
      .catch((res) => {
        if ("messages" in res.response.data) return res.response.data;

        const messages: Messages = {
          messages: [
            {
              validationType: "error",
              description: res.message,
            },
          ],
        };
        return messages;
      });
  }

  async get(id: string): Promise<Document | Messages> {
    return axios
      .get(`${process.env.VUE_APP_API_HOST}/api/documents/${id}`)
      .then((res) => {
        return res.data;
      })
      .catch((res) => {
        if ("messages" in res.response.data) return res.response.data;

        const messages: Messages = {
          messages: [
            {
              validationType: "error",
              description: res.message,
            },
          ],
        };
        return messages;
      });
  }

  async post(document: Document): Promise<null | Messages> {
    return axios
      .post(`${process.env.VUE_APP_API_HOST}/api/documents`, {
        name: document.name,
        identityNumber: document.identityNumber.replace(/[^a-zA-Z0-9]/g, ""),
      })
      .then((res) => {
        return null;
      })
      .catch((res) => {
        if ("messages" in res.response.data) return res.response.data;

        const messages: Messages = {
          messages: [
            {
              validationType: "error",
              description: res.message,
            },
          ],
        };
        return messages;
      });
  }

  async put(document: Document): Promise<null | Messages> {
    return axios
      .put(`${process.env.VUE_APP_API_HOST}/api/documents/${document.uuid}`, {
        name: document.name,
        identityNumber: document.identityNumber.replace(/[^a-zA-Z0-9]/g, ""),
      })
      .then((res) => {
        return null;
      })
      .catch((res) => {
        if ("messages" in res.response.data) return res.response.data;

        const messages: Messages = {
          messages: [
            {
              validationType: "error",
              description: res.message,
            },
          ],
        };
        return messages;
      });
  }

  async delete(id: string): Promise<null | Messages> {
    return axios
      .delete(`${process.env.VUE_APP_API_HOST}/api/documents/${id}`)
      .then((res) => null)
      .catch((res) => {
        if ("messages" in res.response.data) return res.response.data;

        const messages: Messages = {
          messages: [
            {
              validationType: "error",
              description: res.message,
            },
          ],
        };
        return messages;
      });
  }
}
