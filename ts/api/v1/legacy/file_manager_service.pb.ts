/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../fetch.pb"
import * as LlmarinerFilesServerV1File_manager_service from "../file_manager_service.pb"
export class FilesWorkerService {
  static GetFilePath(req: LlmarinerFilesServerV1File_manager_service.GetFilePathRequest, initReq?: fm.InitReq): Promise<LlmarinerFilesServerV1File_manager_service.GetFilePathResponse> {
    return fm.fetchReq<LlmarinerFilesServerV1File_manager_service.GetFilePathRequest, LlmarinerFilesServerV1File_manager_service.GetFilePathResponse>(`/llmoperator.files.server.v1.FilesWorkerService/GetFilePath`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}