export default interface Messages {
  messages: Message[];
}

interface Message {
  validationType: string;
  description: string;
}

export { Message };
