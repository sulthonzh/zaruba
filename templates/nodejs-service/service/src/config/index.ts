import { RmqEventMap } from "../transport";

export class Config {
    public httpPort: number;
    public serviceName: string;
    public defaultRmqConnectionString: string;
    public rmqEventMap: RmqEventMap;

    constructor() {
        this.httpPort = (process.env.SERVICENAME_HTTP_PORT || 3000) as number;
        this.serviceName = "servicename";
        this.defaultRmqConnectionString = process.env.DEFAULT_RMQ_CONNECTION_STRING || "amqp://localhost:5672/";
        this.rmqEventMap = new RmqEventMap({
            helloRPC: {
                exchangeName: "servicename.helloRPC",
                queueName: "servicename.helloRPC",
            },
            hello: {
                exchangeName: "servicename.helloEvent",
                queueName: "servicename.helloEvent",
            }
        });
    }

}