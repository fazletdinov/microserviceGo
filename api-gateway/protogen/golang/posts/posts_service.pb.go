// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: posts/posts_service.proto

package golang

import (
	reflect "reflect"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_posts_posts_service_proto protoreflect.FileDescriptor

var file_posts_posts_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x6f, 0x73,
	0x74, 0x73, 0x1a, 0x11, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcf, 0x05, 0x0a, 0x0c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x41, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x15, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x45, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x61, 0x70, 0x69, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_posts_posts_service_proto_goTypes = []any{
	(*CreatePostRequest)(nil),     // 0: posts.CreatePostRequest
	(*GetPostRequest)(nil),        // 1: posts.GetPostRequest
	(*GetPostsRequest)(nil),       // 2: posts.GetPostsRequest
	(*UpdatePostRequest)(nil),     // 3: posts.UpdatePostRequest
	(*DeletePostRequest)(nil),     // 4: posts.DeletePostRequest
	(*CreateCommentRequest)(nil),  // 5: posts.CreateCommentRequest
	(*GetCommentsRequest)(nil),    // 6: posts.GetCommentsRequest
	(*GetCommentRequest)(nil),     // 7: posts.GetCommentRequest
	(*DeleteCommentRequest)(nil),  // 8: posts.DeleteCommentRequest
	(*UpdateCommentRequest)(nil),  // 9: posts.UpdateCommentRequest
	(*CreatePostResponse)(nil),    // 10: posts.CreatePostResponse
	(*GetPostResponse)(nil),       // 11: posts.GetPostResponse
	(*GetPostsResponse)(nil),      // 12: posts.GetPostsResponse
	(*UpdatePostResponse)(nil),    // 13: posts.UpdatePostResponse
	(*DeletePostResponse)(nil),    // 14: posts.DeletePostResponse
	(*CreateCommentResponse)(nil), // 15: posts.CreateCommentResponse
	(*GetCommentsResponse)(nil),   // 16: posts.GetCommentsResponse
	(*GetCommentResponse)(nil),    // 17: posts.GetCommentResponse
	(*DeleteCommentResponse)(nil), // 18: posts.DeleteCommentResponse
	(*UpdateCommentResponse)(nil), // 19: posts.UpdateCommentResponse
}
var file_posts_posts_service_proto_depIdxs = []int32{
	0,  // 0: posts.GatewayPosts.CreatePost:input_type -> posts.CreatePostRequest
	1,  // 1: posts.GatewayPosts.GetPostByID:input_type -> posts.GetPostRequest
	2,  // 2: posts.GatewayPosts.GetPosts:input_type -> posts.GetPostsRequest
	3,  // 3: posts.GatewayPosts.UpdatePost:input_type -> posts.UpdatePostRequest
	4,  // 4: posts.GatewayPosts.DeletePost:input_type -> posts.DeletePostRequest
	5,  // 5: posts.GatewayPosts.CreateComment:input_type -> posts.CreateCommentRequest
	6,  // 6: posts.GatewayPosts.GetPostComments:input_type -> posts.GetCommentsRequest
	7,  // 7: posts.GatewayPosts.GetCommentByID:input_type -> posts.GetCommentRequest
	8,  // 8: posts.GatewayPosts.DeletePostComment:input_type -> posts.DeleteCommentRequest
	9,  // 9: posts.GatewayPosts.UpdatePostComment:input_type -> posts.UpdateCommentRequest
	10, // 10: posts.GatewayPosts.CreatePost:output_type -> posts.CreatePostResponse
	11, // 11: posts.GatewayPosts.GetPostByID:output_type -> posts.GetPostResponse
	12, // 12: posts.GatewayPosts.GetPosts:output_type -> posts.GetPostsResponse
	13, // 13: posts.GatewayPosts.UpdatePost:output_type -> posts.UpdatePostResponse
	14, // 14: posts.GatewayPosts.DeletePost:output_type -> posts.DeletePostResponse
	15, // 15: posts.GatewayPosts.CreateComment:output_type -> posts.CreateCommentResponse
	16, // 16: posts.GatewayPosts.GetPostComments:output_type -> posts.GetCommentsResponse
	17, // 17: posts.GatewayPosts.GetCommentByID:output_type -> posts.GetCommentResponse
	18, // 18: posts.GatewayPosts.DeletePostComment:output_type -> posts.DeleteCommentResponse
	19, // 19: posts.GatewayPosts.UpdatePostComment:output_type -> posts.UpdateCommentResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_posts_posts_service_proto_init() }
func file_posts_posts_service_proto_init() {
	if File_posts_posts_service_proto != nil {
		return
	}
	file_posts_posts_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_posts_posts_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_posts_posts_service_proto_goTypes,
		DependencyIndexes: file_posts_posts_service_proto_depIdxs,
	}.Build()
	File_posts_posts_service_proto = out.File
	file_posts_posts_service_proto_rawDesc = nil
	file_posts_posts_service_proto_goTypes = nil
	file_posts_posts_service_proto_depIdxs = nil
}
