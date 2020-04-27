package sql

const Select                            = "SELECT ";
const InteractionCount = "COUNT(Id) frequency ";
const InteractionFields = "Id, Name, CreatedDate, Channel__c, Type__c, Status__c, Operator_Assisted__c, Escalated__c, SBE_Call__c, Gateway_Call__c, Language_Request_Time__c, Start_Time__c, Ring_Time__c, Agent_Answer_Time__c, Agent_Join_Time__c, Active_Time__c, Interpreter_Join_Time__c, End_Time__c, Cancel_Time__c, Wait_Duration_Seconds__c, Customer_Hold_Duration_Seconds__c, Duration_Minutes__c, Notes_Category__c, Notes_Sub_Category__c, Notes__c, Supervisor_Notes__c, Customer_Call_Rating__c, Customer_Call_Rating_Reason__c, Customer_Phone_Number__c, Billable__c, Non_Billable_Reason__c, Conference_ID__c, Customer_Call_ID__c, Meeting_Room__c, Meeting_URL__c, Whiteboard_URL__c, Inhouse_Interpreter_Duration_Seconds__c, Inhouse_Direct_Operator_Duration_Seconds__c, Inhouse_Operator_Duration_Seconds__c, Inhouse_Supervisor_Duration_Seconds__c, OPI_Duration_Seconds__c, Outsource_Duration_Seconds__c, Micro_Call_Center_Duration_Seconds__c, Call_Segments_Duration_Seconds__c, Duration_Seconds__c, Call_Center_Platform__c, Outsource_Interpreter_Name__c";
const LanguageFields = "Language__r.Id, Language__r.Name";
//const ClientFields = "Client__r.Id, Client__r.Name__c, Client__r.Micro_Call_Center__c, Client__r.Physical_Address_State__c";
const InterpreterFields = "Interpreter__r.Id, Interpreter__r.Name";
const AccessCodeFields = "Access_Code__r.Id, Access_Code__r.Type__c";
const AccessCodeDepartmentFields = "Access_Code__r.Department__r.Id, Access_Code__r.Department__r.Name";
const OutsourcePartnerFields = "Outsource_Partner__r.Id,Outsource_Partner__r.Channel__c";
const TasksSubQuery = "(SELECT Id,Name,CreatedDate,Router_Task_ID__c,Router_Assignment_ID__c,Router_Workspace_ID__c,Status__c,Type__c,Priority__c,Escalated__c,Self_Assigned__c,Queue_Time__c,Queue_Timeout_Time__c,Assign_Time__c,Last_Assignment_Timeout_Time__c,Accept_Time__c,Complete_Time__c,Agent__r.Id,Agent__r.Router_Agent_ID__c,Agent__r.Router_Workspace_ID__c,Queue__r.Id,Queue__r.Router_Queue_ID__c,Queue__r.Router_Workspace_ID__c,Queue__r.Name,Queue__r.Type__c,Queue__r.Service__c,Queue__r.User__c FROM Tasks__r)";
const AgentRatingsSubQuery = "(SELECT Id,Name,Rating__c, Agent__r.id, Agent__r.name,Agent__r.Role__c  FROM Agent_Ratings__r order by Survey_End_Time__c desc)";
const EventLogsSubQuery = "(SELECT Id, Event__c, Time__c FROM Event_Logs__r)";
const InhouseCallSegmentsSubQuery = "(SELECT Id, Name, CreatedDate, Start_Time__c, End_Time__c, Call_ID__c, Notes__c, Channel__c, Participant_Mode__c, Agent_Role__c, Direct_Operator__c, Self_Assigned__c, Self_Assign_Time__c, Self_Assign_Language__r.Id, Self_Assign_Language__r.Name, Interaction__r.Id, Interaction__r.Name, Language__r.Id, Language__r.Name, Agent__r.Id, Agent__r.Name, Agent__r.Role__c,Agent__r.Micro_Call_Center__c FROM Inhouse_Call_Segments__r order by End_Time__c desc)";
const OutsourceCallSegmentsSubQuery = "(SELECT Id, Name, CreatedDate, Start_Time__c, End_Time__c, Call_ID__c, Notes__c, Channel__c, Interpreter_Name__c, Interpretation_Start_Time__c, Outsource_Reason__c, Outsource_Partner__r.Id, Outsource_Partner__r.Name, Interaction__r.Id, Interaction__r.Name,Outsource_Language__r.Id,Outsource_Language__r.Name,Language__r.Id,Language__r.Name FROM Outsource_Call_Segments__r order by End_Time__c desc)";
const OpiCallSegmentsSubQuery = "(SELECT Id, Name, CreatedDate, Start_Time__c, End_Time__c, Call_ID__c, Notes__c, Phone_Number__c, Interpreter_Name__c, Interaction__r.Id, Interaction__r.Name, Language__r.Id, Language__r.Name FROM OPI_Call_Segments__r order by End_Time__c desc)";
const DirectCallSegmentsSubQuery = "(SELECT Id, Name, CreatedDate, Start_Time__c, End_Time__c, Call_ID__c, Notes__c, Channel__c, Interaction__r.Id, Interaction__r.Name, Agent__r.Id, Agent__r.Name FROM Direct_Call_Segments__r order by End_Time__c desc)";
const UnknownCallSegmentsSubQuery = "(SELECT Id, Name, CreatedDate, Start_Time__c, End_Time__c, Notes__c, Username__c, Interaction__r.Id, Interaction__r.Name FROM Unknown_Call_Segments__r order by End_Time__c desc)";
const FromInteraction = " FROM Interaction__c"
const ClientFields = "Id, Name, Name__c, Photo__c, Photo_Document_ID__c, CreatedDate, Contact_Phone_Number__c, Contact_Name__c, Contact_Email__c, Status__c, Physical_Address_Line_1__c, Physical_Address_Line_2__c, Physical_Location_City__c, Physical_Address_State__c, Physical_Address_Zip__c, Micro_Call_Center__c, Micro_Call_Center_Router_Workspace_ID__c, Call_Center_Platform__c, Access_Code_Validation_Required__c, PEM_Email__c, PEM_Name__c, PEM_Phone_Number__c, Technical_Support_Number__c";
const FromClient = " FROM ClientX__c"

const SelectQuery = Select + InteractionFields + FromInteraction
const SelectCountQuery = Select + InteractionCount + FromInteraction
const SelectWithAccessCodeQuery = Select + InteractionFields + ", " + AccessCodeFields + ", " + AccessCodeDepartmentFields + FromInteraction
const SelectWithOutsourcePartnerQuery = Select + InteractionFields + ", " + OutsourcePartnerFields + FromInteraction
const SelectWithLanguageQuery = Select + InteractionFields + ", " + LanguageFields + FromInteraction
const SelectWithClientQuery = Select + InteractionFields + ", " + ClientFields + FromInteraction
const SelectWithLanguageAndClientQuery = Select + InteractionFields + ", " + LanguageFields + ", " + ClientFields + FromInteraction
const SelectWithLanguageAndTasksQuery = Select + InteractionFields + ", " + LanguageFields + ", " + TasksSubQuery + FromInteraction
const SelectWithLanguageAndClientAndTasksQuery = Select + InteractionFields + ", " + LanguageFields + ", " + ClientFields + ", " + TasksSubQuery + FromInteraction;
const SelectWithTasksAndEventLogsQuery = Select + InterpreterFields + ", " + TasksSubQuery + ", " + EventLogsSubQuery + FromInteraction;
const SelectWithCallSegmentsQuery = Select + InterpreterFields + ", " + InhouseCallSegmentsSubQuery + ", " + OutsourceCallSegmentsSubQuery + ", " + OpiCallSegmentsSubQuery + ", " + DirectCallSegmentsSubQuery + FromInteraction;
const SelectWithClientAndCallSegmentsQuery = Select + InterpreterFields + ", " + ClientFields + ", " + InhouseCallSegmentsSubQuery + ", " + OutsourceCallSegmentsSubQuery + ", " + OpiCallSegmentsSubQuery + ", " + DirectCallSegmentsSubQuery + FromInteraction;
const SelectWithInterpreterAndOutsourcePartnerAndCallSegmentsQuery = Select + InteractionFields + ", " + InterpreterFields + ", " + OutsourcePartnerFields + ", " + InhouseCallSegmentsSubQuery + ", " + OutsourceCallSegmentsSubQuery + ", " + OpiCallSegmentsSubQuery + ", " + DirectCallSegmentsSubQuery + ", " + UnknownCallSegmentsSubQuery + FromInteraction;
const SelectWithLanguageAndClientAndCallSegmentsQuery = Select + InterpreterFields + ", " + LanguageFields + ", " + ClientFields + ", " + InhouseCallSegmentsSubQuery + ", " + OutsourceCallSegmentsSubQuery + ", " + OpiCallSegmentsSubQuery + ", " + DirectCallSegmentsSubQuery + FromInteraction;
const SelectCallSegmentsQuery = Select + InhouseCallSegmentsSubQuery + ", " + OutsourceCallSegmentsSubQuery + ", " + OpiCallSegmentsSubQuery + ", " + DirectCallSegmentsSubQuery + ", " + UnknownCallSegmentsSubQuery + FromInteraction;
const SelectTasksQuery = Select + TasksSubQuery + FromInteraction;
const EventLogsSubQueryWithOrderBy = "(SELECT Id, Event__c, Time__c FROM Event_Logs__r ORDER BY Time__c)";
const SelectFromCustomerHoldTimeQuery = "SELECT Id,Hold_Start_Time__c,Hold_End_Time__c,Hold_Duration_Seconds__c FROM Customer_Hold_Time__c";
const SelectFromCallSheetTimeQuery = "SELECT Id,Start_Time__c,Close_Time__c FROM Call_Sheet_Time__c";
const SelectFromAfterCallWorkQuery = "SELECT Id,Start_Time__c,End_Time__c FROM After_Call_Work__c";
const SelectFromOutsourceRequestQuery = "SELECT Id,Outsource_Initiate_Time__c,Outsource_Cancel_Time__c,Agent__r.Id,Agent__r.Name,Interaction__r.Id,Interaction__r.Name,Interaction__r.Channel__c,Interaction__r.Type__c,Interaction__r.Meeting_Room__c,Interaction__r.Conference_ID__c,Outsource_Partner__r.Id,Outsource_Partner__r.Channel__c,Outsource_Partner__r.Place_Call_API_URL__c,Outsource_Partner__r.Cancel_Call_API_URL__c,Outsource_Partner__r.Call_Status_API_URL__c,Outsource_Partner__r.API_Authorization_ID__c FROM Outsource_Request__c";
const SelectFromEventLogQuery = "SELECT Id, Event__c, Time__c FROM Event_Log__c";

const Show                             = "Show_";
const Order                            = "Order=";
const Category                         = "Category=";
const FromInteractionWhereIdEquals = " FROM Interaction__c WHERE id = '";
// Standard Fields
const FieldId = "Id";
const FieldName = "Name";
const FieldCreatedDate = "CreatedDate";
const FieldCreatedById = "CreatedById";
const FieldLastModifiedDate = "LastModifiedDate";
const FieldLastModifiedById = "LastModifiedById";
const FieldSystemModstamp = "SystemModstamp";
const FieldLastActivityDate = "LastActivityDate";
const FieldLastViewedDate = "LastViewedDate";
const FieldLastReferencedDate = "LastReferencedDate";
const FieldIsDeleted = "IsDeleted";
const FieldOwnerId = "OwnerId";
// Custom Fields
const FieldType = "Type__c";
const FieldChannel = "Channel__c";
const FieldStatus = "Status__c";
const FieldStartTime = "Start_Time__c";
const FieldActiveTime = "Active_Time__c";
const FieldEndTime = "End_Time__c";
const FieldOperatorAssisted = "Operator_Assisted__c";
const FieldBillable = "Billable__c";
const FieldEscalated = "Escalated__c";
const FieldSbeCall = "SBE_Call__c";
const FieldChatUrl = "Chat_Url__c";
const FieldWhiteboardUrl = "Whiteboard_URL__c";
const FieldMeetingUrl = "Meeting_URL__c";
const FieldMeetingRoom = "Meeting_Room__c";
const FieldConferenceId = "Conference_ID__c";
const FieldCustomerCallId = "Customer_Call_ID__c";
const FieldCustomerPhoneNumber = "Customer_Phone_Number__c";
const FieldCustomerCallRating = "Customer_Call_Rating__c";
const FieldCustomerCallRatingReason = "Customer_Call_Rating_Reason__c";
const FieldCustomerHoldDurationSeconds = "Customer_Hold_Duration_Seconds__c";
const FieldWaitDurationSeconds = "Wait_Duration_Seconds__c";
const FieldAgentAnswerDurationSeconds = "Agent_Answer_Duration_Seconds__c";
const FieldDurationMinutes = "Duration_Minutes__c";
const FieldDurationSeconds = "Duration_Seconds__c";
const FieldNotes = "Notes__c";
const FieldNotesCategory = "Notes_Category__c";
const FieldNotesSubCategory = "Notes_Sub_Category__c";
const FieldSupervisorNotes = "Supervisor_Notes__c";
const FieldLanguage = "Language__c";
const FieldMartti = "MARTTI__c";
const FieldInterpreter = "Interpreter__c";
const FieldDepartment = "Department__c";
const FieldClient = "Client__c";
const FieldAccessCode = "Access_Code__c";
const FieldAccessCodeCode = "Code__c";
const FieldClientName = "Name__c";
const FieldEventLogEvent = "Event__c";
const FieldEventLogTime = "Time__c";
const FieldAgentRatingRating = "Rating__c";
const FieldCallratingNotes = "Notes__c";
const FieldAgentRole = "Role__c";

const SobjectInteraction = "Interaction__c";
const SobjectAccessCode = "Access_CodeX__c";

const ParentLanguage = "Language__r";
const ParentMartti = "MARTTI__r";
const ParentInterpreter = "Interpreter__r";
const ParentDepartment = "Department__r";
const ParentClient = "Client__r";
const ParentAgent = "Agent__r";
const ParentAccessCode = "Access_Code__r";

const ChildEventLogs = "Event_Logs__r";
const ChildAgentRatings = "Agent_Ratings__r";
const ChildInhouseCallSegments = "Inhouse_Call_Segments__r";
const ChildOutsourceCallSegments = "Outsource_Call_Segments__r";
const ChildOpiCallSegments = "OPI_Call_Segments__r";

const InteractionId = "Interaction_Id__c";
const InteractionName = "Interaction_Name__c";
const AccessCodeRelation = "Access_Code__r.";

const ParentFields = "Language__r.Name, Interpreter__r.Name, Client__r.Id, Client__r.Name__c, MARTTI__r.Name, MARTTI__r.Department__r.Name, Department__r.Id, Department__r.Name, Access_Code__r.Code__c"
//private static final List<String> STANDARD_FIELDS = Arrays.asList(FIELD_ID, FIELD_NAME, FIELD_TYPE, FIELD_CHANNEL,
//FIELD_STATUS, FIELD_START_TIME, FIELD_ACTIVE_TIME, FIELD_END_TIME, FIELD_OPERATOR_ASSISTED, FIELD_BILLABLE,
//FIELD_ESCALATED, FIELD_SBE_CALL, FIELD_CHAT_URL, FIELD_WHITEBOARD_URL, FIELD_MEETING_URL,
//FIELD_MEETING_ROOM, FIELD_CONFERENCE_ID, FIELD_CUSTOMER_CALL_ID, FIELD_CUSTOMER_PHONE_NUMBER,
//FIELD_CUSTOMER_CALL_RATING, FIELD_CUSTOMER_CALL_RATING_REASON, FIELD_CUSTOMER_HOLD_DURATION_SECONDS,
//FIELD_WAIT_DURATION_SECONDS, FIELD_AGENT_ANSWER_DURATION_SECONDS, FIELD_DURATION_MINUTES,
//FIELD_DURATION_SECONDS, FIELD_NOTES, FIELD_NOTES_CATEGORY, FIELD_NOTES_SUB_CATEGORY, FIELD_SUPERVISOR_NOTES,
//FIELD_LANGUAGE, FIELD_MARTTI, FIELD_INTERPRETER, FIELD_DEPARTMENT, FIELD_CLIENT, FIELD_ACCESS_CODE);
//const Select = "SELECT "
//const InteractionFields = "Id, Name, CreatedDate, Channel__c, Type__c, Status__c, Operator_Assisted__c, Escalated__c, SBE_Call__c, Gateway_Call__c, Language_Request_Time__c, Start_Time__c, Ring_Time__c, Agent_Answer_Time__c, Agent_Join_Time__c, Active_Time__c, Interpreter_Join_Time__c, End_Time__c, Cancel_Time__c, Wait_Duration_Seconds__c, Customer_Hold_Duration_Seconds__c, Duration_Minutes__c, Notes_Category__c, Notes_Sub_Category__c, Notes__c, Supervisor_Notes__c, Customer_Call_Rating__c, Customer_Call_Rating_Reason__c, Customer_Phone_Number__c, Billable__c, Non_Billable_Reason__c, Conference_ID__c, Customer_Call_ID__c, Meeting_Room__c, Meeting_URL__c, Whiteboard_URL__c, Inhouse_Interpreter_Duration_Seconds__c, Inhouse_Direct_Operator_Duration_Seconds__c, Inhouse_Operator_Duration_Seconds__c, Inhouse_Supervisor_Duration_Seconds__c, OPI_Duration_Seconds__c, Outsource_Duration_Seconds__c, Micro_Call_Center_Duration_Seconds__c, Call_Segments_Duration_Seconds__c, Duration_Seconds__c, Call_Center_Platform__c, Outsource_Interpreter_Name__c"
//const FromInteraction = " FROM Interaction__c";

//const SelectQuery = Select + InteractionFields + FromInteraction


