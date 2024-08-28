import * as fm from "../../fetch.pb";
export type File = {
    id?: string;
    bytes?: string;
    createdAt?: string;
    filename?: string;
    object?: string;
    purpose?: string;
};
export type ListFilesRequest = {
    purpose?: string;
};
export type ListFilesResponse = {
    object?: string;
    data?: File[];
};
export type GetFileRequest = {
    id?: string;
};
export type DeleteFileRequest = {
    id?: string;
};
export type DeleteFileResponse = {
    id?: string;
    object?: string;
    deleted?: boolean;
};
export type GetFilePathRequest = {
    id?: string;
};
export type GetFilePathResponse = {
    path?: string;
    filename?: string;
};
export declare class FilesService {
    static ListFiles(req: ListFilesRequest, initReq?: fm.InitReq): Promise<ListFilesResponse>;
    static GetFile(req: GetFileRequest, initReq?: fm.InitReq): Promise<File>;
    static DeleteFile(req: DeleteFileRequest, initReq?: fm.InitReq): Promise<DeleteFileResponse>;
}
export declare class FilesWorkerService {
    static GetFilePath(req: GetFilePathRequest, initReq?: fm.InitReq): Promise<GetFilePathResponse>;
}
export declare class FilesInternalService {
    static GetFilePath(req: GetFilePathRequest, initReq?: fm.InitReq): Promise<GetFilePathResponse>;
}
