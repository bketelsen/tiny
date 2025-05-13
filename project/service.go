package project

import "strings"

type Service struct {
	Name        string
	Description string
	EndpointMap map[string]*Endpoint
	MessageMap  map[string]*Message
	EnumMap     map[string]*Enum
}

func (s *Service) GetEndpoint(name string) (*Endpoint, bool) {
	endpoint, ok := s.EndpointMap[name]
	return endpoint, ok
}

func (s *Service) GetMessage(name string) (*Message, bool) {
	message, ok := s.MessageMap[name]
	return message, ok
}

func (s *Service) GetEnum(name string) (*Enum, bool) {
	enum, ok := s.EnumMap[name]
	return enum, ok
}

func (s *Service) GetAllEndpoints() []*Endpoint {
	endpoints := make([]*Endpoint, 0, len(s.EndpointMap))
	for _, endpoint := range s.EndpointMap {
		endpoints = append(endpoints, endpoint)
	}
	return endpoints
}

func (s *Service) GetAllMessages() []*Message {
	messages := make([]*Message, 0, len(s.MessageMap))
	for _, message := range s.MessageMap {
		messages = append(messages, message)
	}
	return messages
}

func (s *Service) GetAllEnums() []*Enum {
	enums := make([]*Enum, 0, len(s.EnumMap))
	for _, enum := range s.EnumMap {
		enums = append(enums, enum)
	}
	return enums
}

func (s *Service) DirectoryName() string {
	if s == nil {
		return ""
	}
	return strings.ToLower(s.Name)
}
