#ifndef STATMENT_HANDLER_H
#define STATMENT_HANDLER_H
#include<string>
enum PrepareResult {
    PREPARE_SUCCESS,
    PREPARE_UNRECOGNIZED_STATEMENT
};

enum StatementType {
    STATEMENT_INSERT,
    STATEMENT_SELECT
};

struct Statement {
    StatementType type;
};

PrepareResult prepare_statement(std::string &input_buffer, Statement &statement); 
void execute_statement(Statement &statement);

#endif
