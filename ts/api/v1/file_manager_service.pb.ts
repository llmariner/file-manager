/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
export type File = {
  id?: string
  bytes?: string
  created_at?: string
  filename?: string
  object?: string
  purpose?: string
  object_store_path?: string
}

export type ListFilesRequest = {
  purpose?: string
  after?: string
  limit?: number
  order?: string
}

export type ListFilesResponse = {
  object?: string
  data?: File[]
  has_more?: boolean
  total_items?: number
}

export type GetFileRequest = {
  id?: string
}

export type DeleteFileRequest = {
  id?: string
}

export type DeleteFileResponse = {
  id?: string
  object?: string
  deleted?: boolean
}

export type CreateFileFromObjectPathRequest = {
  object_path?: string
  purpose?: string
}

export type GetFilePathRequest = {
  id?: string
}

export type GetFilePathResponse = {
  path?: string
  filename?: string
}

export class FilesService {
  static ListFiles(req: ListFilesRequest, initReq?: fm.InitReq): Promise<ListFilesResponse> {
    return fm.fetchReq<ListFilesRequest, ListFilesResponse>(`/v1/files?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetFile(req: GetFileRequest, initReq?: fm.InitReq): Promise<File> {
    return fm.fetchReq<GetFileRequest, File>(`/v1/files/${req["id"]}?${fm.renderURLSearchParams(req, ["id"])}`, {...initReq, method: "GET"})
  }
  static DeleteFile(req: DeleteFileRequest, initReq?: fm.InitReq): Promise<DeleteFileResponse> {
    return fm.fetchReq<DeleteFileRequest, DeleteFileResponse>(`/v1/files/${req["id"]}`, {...initReq, method: "DELETE"})
  }
  static CreateFileFromObjectPath(req: CreateFileFromObjectPathRequest, initReq?: fm.InitReq): Promise<File> {
    return fm.fetchReq<CreateFileFromObjectPathRequest, File>(`/v1/files:createFromObjectPath`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}
export class FilesWorkerService {
  static GetFilePath(req: GetFilePathRequest, initReq?: fm.InitReq): Promise<GetFilePathResponse> {
    return fm.fetchReq<GetFilePathRequest, GetFilePathResponse>(`/llmariner.files.server.v1.FilesWorkerService/GetFilePath`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}
export class FilesInternalService {
  static GetFilePath(req: GetFilePathRequest, initReq?: fm.InitReq): Promise<GetFilePathResponse> {
    return fm.fetchReq<GetFilePathRequest, GetFilePathResponse>(`/llmariner.files.server.v1.FilesInternalService/GetFilePath`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}