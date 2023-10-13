import type { ApiSuccess } from "#/api";

export const apiLog = {
  async getLog() {
    const { data } = await axios.get<string>('api/v1/log');
    
    // pretty print
    return data.replaceAll(`"level":"warn"`, "WARN]").replaceAll(`"level":"error"`, "ERROR]").
    replaceAll(`"level":"info"`, "INFO]").replaceAll(`"level":"debug"`, "DEBUG]").replaceAll("{", "[").replaceAll("}", "").replaceAll(`,"`, " ").replaceAll('"', "").replaceAll(`time:`, "[").replaceAll(` message:`, "] ");
  },
  
  async clearLog() {
    const { data } = await axios.get<ApiSuccess>('api/v1/log/clear');
    return data;
  },
};
