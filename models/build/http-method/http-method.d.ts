export declare const protobufPackage = "httpmethod";
export declare enum Method {
    GET = 0,
    POST = 1,
    PUT = 2,
    DELETE = 3,
    UNRECOGNIZED = -1
}
export declare function methodFromJSON(object: any): Method;
export declare function methodToJSON(object: Method): string;
