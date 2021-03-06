#!/usr/bin/env python
# -*- encoding:utf8 -*-

# 这个文件是根据项目来更改的文件 

import sys
import re
from string import strip

from panpbtool.conf import protolabel
from panpbtool.data.baserule import baserule

class customrule(baserule):
    context = None
    def __init__(self, context):
        self.context = context
        fp = open(context.write_path + "/protomsg.go", "r")
        id_descript_data = fp.read()
        fp.close()
        context.id_descript_data = id_descript_data

    def get_namespace(self):
        context = self.context
        raw_data = context.raw_data
        if not hasattr(context, "namespace"):
            project_namespace = ""
            for proto_file in raw_data.proto_file:
                project_namespace = proto_file.package
                index = project_namespace.rfind('.')
                if index != -1:
                    project_namespace = project_namespace[0:index]
                break
            context.namespace = project_namespace
        return context.namespace

    def get_module_name(self, module_desc):
        namespace = self.get_namespace()
        name = module_desc.package[(len(namespace)+1):]
        if len(name) == 0:
            return "global"
        return name

    def is_protocol(self, message_desc):
        message_name = message_desc.name

        if message_name.startswith('C2S'):
            return True
            
        if message_name.startswith('S2C'):
            return True

        return False

    def get_protocol_id(self, message_desc):
        context = self.context
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

    def get_protocol_category(self, message_desc):
        message_name = message_desc.name

        if message_name.startswith('C2S'):
            return "Request"

        if message_name.startswith('S2C'):
            return "Response"

        return "Notification"