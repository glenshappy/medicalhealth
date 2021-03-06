package cast

// Code generated by cdproto-gen. DO NOT EDIT.

// EventSinksUpdated this is fired whenever the list of available sinks
// changes. A sink is a device or a software surface that you can cast to.
type EventSinksUpdated struct {
	SinkNames []string `json:"sinkNames"`
}

// EventIssueUpdated this is fired whenever the outstanding issue/error
// message changes. |issueMessage| is empty if there is no issue.
type EventIssueUpdated struct {
	IssueMessage string `json:"issueMessage"`
}
