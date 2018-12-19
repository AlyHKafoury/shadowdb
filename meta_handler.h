#ifndef META_COMMAND_H
#define META_COMMAND_H
#include<string>

enum MetaCommandResult {
    META_COMMAND_SUCCESS,
    META_COMMAND_UNRECOGNIZED_COMMAND
};

MetaCommandResult do_meta_command(std::string &input_buffer);
#endif
