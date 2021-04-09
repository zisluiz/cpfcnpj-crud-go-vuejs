import Messages from "@/model/Messages";
import Document from "@/model/Document";
import axios from "axios";

export default class BlockListService {
  async blockDocuments(
    documents: Document[],
    block: boolean
  ): Promise<null | Messages> {
    return axios
      .put(
        `${process.env.VUE_APP_API_HOST}/api/${
          block ? "blocklist" : "unblocklist"
        }`,
        {
          uuids: documents.map((document) => document.uuid),
        }
      )
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
}
