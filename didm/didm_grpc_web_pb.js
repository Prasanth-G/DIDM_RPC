/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = require('./didm_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.DownloadServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.DownloadServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.DistributedDownloadRequest,
 *   !proto.DownloadResponse>}
 */
const methodInfo_DownloadService_DistributeDownload = new grpc.web.AbstractClientBase.MethodInfo(
  proto.DownloadResponse,
  /** @param {!proto.DistributedDownloadRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.DownloadResponse.deserializeBinary
);


/**
 * @param {!proto.DistributedDownloadRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.DownloadResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.DownloadResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.DownloadServiceClient.prototype.distributeDownload =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/DownloadService/DistributeDownload',
      request,
      metadata || {},
      methodInfo_DownloadService_DistributeDownload,
      callback);
};


/**
 * @param {!proto.DistributedDownloadRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.DownloadResponse>}
 *     A native promise that resolves to the response
 */
proto.DownloadServicePromiseClient.prototype.distributeDownload =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/DownloadService/DistributeDownload',
      request,
      metadata || {},
      methodInfo_DownloadService_DistributeDownload);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.DownloadRequest,
 *   !proto.DownloadResponse>}
 */
const methodInfo_DownloadService_Download = new grpc.web.AbstractClientBase.MethodInfo(
  proto.DownloadResponse,
  /** @param {!proto.DownloadRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.DownloadResponse.deserializeBinary
);


/**
 * @param {!proto.DownloadRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.DownloadResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.DownloadResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.DownloadServiceClient.prototype.download =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/DownloadService/Download',
      request,
      metadata || {},
      methodInfo_DownloadService_Download,
      callback);
};


/**
 * @param {!proto.DownloadRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.DownloadResponse>}
 *     A native promise that resolves to the response
 */
proto.DownloadServicePromiseClient.prototype.download =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/DownloadService/Download',
      request,
      metadata || {},
      methodInfo_DownloadService_Download);
};


module.exports = proto;

