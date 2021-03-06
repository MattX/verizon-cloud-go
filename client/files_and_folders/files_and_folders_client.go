// Code generated by go-swagger; DO NOT EDIT.

package files_and_folders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new files and folders API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for files and folders API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateVirtualFolder(params *CreateVirtualFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateVirtualFolderCreated, error)

	DeleteDelete(params *DeleteDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDeleteOK, error)

	DeleteVirtualFolder(params *DeleteVirtualFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteVirtualFolderNoContent, error)

	GetFiles(params *GetFilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFilesOK, error)

	GetFullview(params *GetFullviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFullviewOK, *GetFullviewResetContent, error)

	GetMetadata(params *GetMetadataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetMetadataOK, error)

	GetMetadataPath(params *GetMetadataPathParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetMetadataPathOK, error)

	GetSearch(params *GetSearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSearchOK, error)

	GetThumbnails(params *GetThumbnailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetThumbnailsOK, error)

	PostCopy(params *PostCopyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostCopyOK, error)

	PostCreatefolder(params *PostCreatefolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostCreatefolderOK, *PostCreatefolderCreated, error)

	PostMove(params *PostMoveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostMoveOK, error)

	PostRename(params *PostRenameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostRenameOK, error)

	PurgeTrash(params *PurgeTrashParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PurgeTrashOK, error)

	RenameVirtualFolder(params *RenameVirtualFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RenameVirtualFolderOK, error)

	RestoreTrash(params *RestoreTrashParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestoreTrashNoContent, error)

	RetrieveTrash(params *RetrieveTrashParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveTrashOK, error)

	UploadFileCreate(params *UploadFileCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadFileCreateCreated, error)

	UploadFileIntent(params *UploadFileIntentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadFileIntentOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateVirtualFolder creates a virtual folder

  Create a virtual folder for the user.
*/
func (a *Client) CreateVirtualFolder(params *CreateVirtualFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateVirtualFolderCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVirtualFolderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create_virtual_folder",
		Method:             "POST",
		PathPattern:        "/virtualfolder/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateVirtualFolderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateVirtualFolderCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create_virtual_folder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteDelete deletes a file folder

  Delete a file/folder at the specified path.
*/
func (a *Client) DeleteDelete(params *DeleteDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_delete",
		Method:             "DELETE",
		PathPattern:        "/fops/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteVirtualFolder deletes a virtual folder

  Delete a virtual folder for the user.
*/
func (a *Client) DeleteVirtualFolder(params *DeleteVirtualFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteVirtualFolderNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVirtualFolderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_virtual_folder",
		Method:             "DELETE",
		PathPattern:        "/virtualfolder",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVirtualFolderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVirtualFolderNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_virtual_folder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFiles retrieves file content

  Returns the content of a file at the specified path.
*/
func (a *Client) GetFiles(params *GetFilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_files",
		Method:             "GET",
		PathPattern:        "/files/{path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_files: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFullview retrieves information on all file and folder content
*/
func (a *Client) GetFullview(params *GetFullviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFullviewOK, *GetFullviewResetContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFullviewParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_fullview",
		Method:             "GET",
		PathPattern:        "/fullview",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFullviewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetFullviewOK:
		return value, nil, nil
	case *GetFullviewResetContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for files_and_folders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetMetadata gets metadata for the root folder contents

  Fetch metadata for all top-level folders.
*/
func (a *Client) GetMetadata(params *GetMetadataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMetadataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_metadata",
		Method:             "GET",
		PathPattern:        "/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMetadataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_metadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetMetadataPath gets metadata for a file or folder in a user s repository

  Fetch metadata for the file or folder under the specified path.
*/
func (a *Client) GetMetadataPath(params *GetMetadataPathParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetMetadataPathOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMetadataPathParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_metadata_path",
		Method:             "GET",
		PathPattern:        "/metadata/{path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMetadataPathReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMetadataPathOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_metadata_path: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSearch searches for files and folders

  Search for files and folders. Search queries can return objects containing words across all metadata fields or just specific fields
*/
func (a *Client) GetSearch(params *GetSearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_search",
		Method:             "GET",
		PathPattern:        "/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetThumbnails retrieves thumbnails for a file
*/
func (a *Client) GetThumbnails(params *GetThumbnailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetThumbnailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetThumbnailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_thumbnails",
		Method:             "GET",
		PathPattern:        "/thumbnails/{content-token}",
		ProducesMediaTypes: []string{"image/jpeg"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetThumbnailsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetThumbnailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_thumbnails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostCopy creates a copy of a file folder

  Create a copy of a file or folder. Any missing folder in the target path will be created.
*/
func (a *Client) PostCopy(params *PostCopyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostCopyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCopyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_copy",
		Method:             "POST",
		PathPattern:        "/fops/copy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCopyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCopyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_copy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostCreatefolder creates a folder

  Create a folder at the specified path.
*/
func (a *Client) PostCreatefolder(params *PostCreatefolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostCreatefolderOK, *PostCreatefolderCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCreatefolderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_createfolder",
		Method:             "POST",
		PathPattern:        "/fops/createfolder",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCreatefolderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostCreatefolderOK:
		return value, nil, nil
	case *PostCreatefolderCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for files_and_folders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostMove moves a file or folder

  Moves a file or folder. Any missing folder in the target path will be created.
*/
func (a *Client) PostMove(params *PostMoveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostMoveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostMoveParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_move",
		Method:             "POST",
		PathPattern:        "/fops/move",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostMoveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostMoveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_move: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostRename renames a file folder

  Rename a file or folder.
*/
func (a *Client) PostRename(params *PostRenameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostRenameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRenameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_rename",
		Method:             "POST",
		PathPattern:        "/fops/rename",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostRenameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostRenameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_rename: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PurgeTrash purges all the files and folder from a repository

  Purges all the files and folder that are in a deleted state in the repository. Irrevocably deletes all paths that are marked as deleted in a repository at the time the call is made.
*/
func (a *Client) PurgeTrash(params *PurgeTrashParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PurgeTrashOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPurgeTrashParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "purge_trash",
		Method:             "DELETE",
		PathPattern:        "/trash",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PurgeTrashReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PurgeTrashOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for purge_trash: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RenameVirtualFolder renames a virtual folder

  Rename a virtual folder for the user.
*/
func (a *Client) RenameVirtualFolder(params *RenameVirtualFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RenameVirtualFolderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRenameVirtualFolderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "rename_virtual_folder",
		Method:             "PUT",
		PathPattern:        "/virtualfolder",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RenameVirtualFolderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RenameVirtualFolderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for rename_virtual_folder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RestoreTrash restores files or folders

  Restores files or folders that are marked as deleted. When called on a folder, the API creates a new version of the folder and any child folder, and it restores all files in these folders.
*/
func (a *Client) RestoreTrash(params *RestoreTrashParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestoreTrashNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreTrashParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "restore_trash",
		Method:             "POST",
		PathPattern:        "/fops/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RestoreTrashReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestoreTrashNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for restore_trash: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrieveTrash retrieves a list of deleted files and folders from a single user repository

  Retrieves a list of files and folders that have been deleted from a single user repository.
*/
func (a *Client) RetrieveTrash(params *RetrieveTrashParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveTrashOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveTrashParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "retrieve_trash",
		Method:             "GET",
		PathPattern:        "/trash",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveTrashReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RetrieveTrashOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for retrieve_trash: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UploadFileCreate completes chunk file upload

  Indicates that file content upload is complete, and the file should be created.
*/
func (a *Client) UploadFileCreate(params *UploadFileCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadFileCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadFileCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "upload_file_create",
		Method:             "POST",
		PathPattern:        "/commit/{uploadid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UploadFileCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UploadFileCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for upload_file_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UploadFileIntent initiates a file upload intent request

  Initiates a file upload intent request. Returns file upload URL and file commit URL in case of file is uploaded in chunk
*/
func (a *Client) UploadFileIntent(params *UploadFileIntentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadFileIntentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadFileIntentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "upload_file_intent",
		Method:             "GET",
		PathPattern:        "/fileupload/intent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UploadFileIntentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UploadFileIntentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for upload_file_intent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
