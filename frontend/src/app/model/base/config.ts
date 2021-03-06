export interface IConfig {
  app: {
    name: string;
    language: string;
  };
  apiServer: {
    endpoint: string;
    websocket: string;
    auth: {
      type: string;
      session: string;
      token: string;
      jwt: string;
    }
    prefix: {
      quant: string;
      time: string;
      system: string;
    }
  }
}
