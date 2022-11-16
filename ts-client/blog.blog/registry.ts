import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreatePost } from "./types/blog/blog/tx";
import { MsgUploadW3S } from "./types/blog/blog/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/blog.blog.MsgCreatePost", MsgCreatePost],
    ["/blog.blog.MsgUploadW3S", MsgUploadW3S],
    
];

export { msgTypes }