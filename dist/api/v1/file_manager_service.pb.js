/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/
import * as fm from "../../fetch.pb";
export class FilesService {
    static ListFiles(req, initReq) {
        return fm.fetchReq(`/v1/files?${fm.renderURLSearchParams(req, [])}`, Object.assign(Object.assign({}, initReq), { method: "GET" }));
    }
    static GetFile(req, initReq) {
        return fm.fetchReq(`/v1/files/${req["id"]}?${fm.renderURLSearchParams(req, ["id"])}`, Object.assign(Object.assign({}, initReq), { method: "GET" }));
    }
    static DeleteFile(req, initReq) {
        return fm.fetchReq(`/v1/files/${req["id"]}`, Object.assign(Object.assign({}, initReq), { method: "DELETE" }));
    }
}
export class FilesWorkerService {
    static GetFilePath(req, initReq) {
        return fm.fetchReq(`/llmariner.files.server.v1.FilesWorkerService/GetFilePath`, Object.assign(Object.assign({}, initReq), { method: "POST", body: JSON.stringify(req) }));
    }
}
export class FilesInternalService {
    static GetFilePath(req, initReq) {
        return fm.fetchReq(`/llmariner.files.server.v1.FilesInternalService/GetFilePath`, Object.assign(Object.assign({}, initReq), { method: "POST", body: JSON.stringify(req) }));
    }
}
