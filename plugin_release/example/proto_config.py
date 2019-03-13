#!/usr/bin/env python
# -*- encoding:utf8 -*-

# 这个文件是根据项目来更改的文件 

import sys
import re
from panpbtool.conf import protolabel
from string import strip

import google.protobuf.compiler.plugin_pb2 as plugin_pb2
import google.protobuf.descriptor_pb2 as descriptor_pb2

FDP = descriptor_pb2.FieldDescriptorProto


# Windows will mangle our line-endings unless we do this.
def fix_line_ending():
    if sys.platform == "win32":
        import os
        import msvcrt  # pylint: disable=import-error

        msvcrt.setmode(sys.stdout.fileno(),
                       os.O_BINARY)  # pylint: disable=E1103 ; the Windows version of os has O_BINARY
        msvcrt.setmode(sys.stderr.fileno(),
                       os.O_BINARY)  # pylint: disable=E1103 ; the Windows version of os has O_BINARY
        msvcrt.setmode(sys.stdin.fileno(),
                       os.O_BINARY)  # pylint: disable=E1103 ; the Windows version of os has O_BINARY

fix_line_ending()

def init(context):
    fp = open("test.bin", 'rb')
    plugin_require_bin = fp.read()
    fp.close()

    fp = open("protomsg.go", "r")
    id_descript_data = fp.read()
    fp.close()
    
    code_gen_req = plugin_pb2.CodeGeneratorRequest()
    code_gen_req.ParseFromString(plugin_require_bin)
    
    context.raw_data = code_gen_req
    context.id_descript_data = id_descript_data

    if not hasattr(context, "namespace"):
        project_namespace = ""
        for proto_file in context.raw_data.proto_file:
            project_namespace = proto_file.package
            index = project_namespace.rfind('.')
            if index != -1:
                project_namespace = project_namespace[0:index]
            break
        context.namespace = project_namespace
    
def get_namespace(context):
    return context.namespace

def get_module_name(context, module_desc):
    name = module_desc.package[(len(context.namespace)+1):]
    if len(name) == 0:
        return "globle"
    return name

def is_protocol(context, message_desc):
    message_name = message_desc.name

    if message_name.startswith('C2S'):
        return True
        
    if message_name.startswith('S2C'):
        return True

    return False

def get_protocol_id(context, message_desc):
    if not hasattr(context, "proto_ids"):
        id_descript_data = context.id_descript_data
        digi_str = re.findall(r'.*=.*', id_descript_data, re.MULTILINE)
        proto_ids = {}
        for _str in digi_str:
            _comment = ''
            _ret = re.findall(r'\/\/[^\n]*', _str, re.MULTILINE)
            if len(_ret) > 0:
                _str = _str.replace(_ret[0], '')
                _comment = _ret[0]
            _ret = re.findall(r'(.*)=(.*)', _str, re.MULTILINE)
            if len(_ret[0]) > 1:
                _name = strip(_ret[0][0]).replace('_', '').lower()
                _id = strip(_ret[0][1])
                proto_ids[_name + 'proto'] = {
                    'name': _name,
                    'id': _id,
                    'comment': _comment,
                }
        context.proto_ids = proto_ids

    message_name = message_desc.name.lower()
    if message_name in context.proto_ids:
        id = context.proto_ids[message_name]['id']
        if id is not None:
            return int(id)
    else:
        print(message_name + " can not find id/n")
        return -1

    return -1

def get_protocol_category(context, message_desc):
    message_name = message_desc.name

    if message_name.startswith('C2S'):
        return "Request"

    if message_name.startswith('S2C'):
        return "Response"

    return "Notification"