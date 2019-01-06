#include"meta_handler.h"
#include<string>
#include<iostream>
using namespace std;

MetaCommandResult do_meta_command(string &input_buffer) {
    if(input_buffer == ".exit"){
        exit(EXIT_SUCCESS);
    }else {
        return META_COMMAND_UNRECOGNIZED_COMMAND;
    }
}


