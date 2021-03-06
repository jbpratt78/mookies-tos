// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: schema/tos.proto
//
// For information on using the generated types, please see the documentation:
//   https://github.com/apple/swift-protobuf/

import Foundation
import SwiftProtobuf

// If the compiler emits an error on this type, it is because this file
// was generated by a version of the `protoc` Swift plug-in that is
// incompatible with the version of SwiftProtobuf to which you are linking.
// Please ensure that you are building against the same version of the API
// that was used to generate this file.
private struct _GeneratedWithProtocGenSwiftVersion: SwiftProtobuf.ProtobufAPIVersionCheck {
    struct _2: SwiftProtobuf.ProtobufAPIVersion_2 {}
    typealias Version = _2
}

public struct Tospb_Item {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    public var name = String()

    public var id: Int64 = 0

    public var price: Float = 0

    /// @inject_tag: db:"category_id"
    public var categoryID: Int64 = 0

    /// @inject_tag: db:"order_item_id"
    public var orderItemID: Int64 = 0

    public var options: [Tospb_Option] = []

    public var unknownFields = SwiftProtobuf.UnknownStorage()

    public init() {}
}

public struct Tospb_Option {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    public var name = String()

    public var id: Int64 = 0

    public var price: Float = 0

    public var selected: Bool = false

    public var unknownFields = SwiftProtobuf.UnknownStorage()

    public init() {}
}

public struct Tospb_Category {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    public var name = String()

    public var id: Int64 = 0

    public var items: [Tospb_Item] = []

    public var unknownFields = SwiftProtobuf.UnknownStorage()

    public init() {}
}

public struct Tospb_Menu {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    public var categories: [Tospb_Category] = []

    public var unknownFields = SwiftProtobuf.UnknownStorage()

    public init() {}
}

public struct Tospb_Order {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    public var id: Int64 = 0

    public var name = String()

    public var items: [Tospb_Item] = []

    public var total: Float = 0

    public var status = String()

    /// @inject_tag: db:"time_ordered"
    public var timeOrdered = String()

    /// @inject_tag: db:"time_complete"
    public var timeComplete = String()

    public var unknownFields = SwiftProtobuf.UnknownStorage()

    public init() {}
}

public struct Tospb_Response {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    public var response = String()

    public var unknownFields = SwiftProtobuf.UnknownStorage()

    public init() {}
}

public struct Tospb_CreateMenuItemResponse {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    public var id: Int64 = 0

    public var unknownFields = SwiftProtobuf.UnknownStorage()

    public init() {}
}

public struct Tospb_CompleteOrderRequest {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    public var id: Int64 = 0

    public var unknownFields = SwiftProtobuf.UnknownStorage()

    public init() {}
}

public struct Tospb_OrdersRequest {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    public var request = String()

    public var unknownFields = SwiftProtobuf.UnknownStorage()

    public init() {}
}

public struct Tospb_OrdersResponse {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    public var orders: [Tospb_Order] = []

    public var unknownFields = SwiftProtobuf.UnknownStorage()

    public init() {}
}

public struct Tospb_Empty {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    public var unknownFields = SwiftProtobuf.UnknownStorage()

    public init() {}
}

public struct Tospb_DeleteMenuItemRequest {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    public var id: Int64 = 0

    public var unknownFields = SwiftProtobuf.UnknownStorage()

    public init() {}
}

public struct Tospb_Ping {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    public var message = String()

    public var delaySeconds: Float = 0

    public var unknownFields = SwiftProtobuf.UnknownStorage()

    public init() {}
}

public struct Tospb_Pong {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    public var message = String()

    public var unknownFields = SwiftProtobuf.UnknownStorage()

    public init() {}
}

// MARK: - Code below here is support for the SwiftProtobuf runtime.

private let _protobuf_package = "tospb"

extension Tospb_Item: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
    public static let protoMessageName: String = _protobuf_package + ".Item"
    public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
        1: .same(proto: "name"),
        2: .same(proto: "id"),
        3: .same(proto: "price"),
        4: .same(proto: "categoryID"),
        5: .same(proto: "orderItemID"),
        6: .same(proto: "options"),
    ]

    public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
        while let fieldNumber = try decoder.nextFieldNumber() {
            // The use of inline closures is to circumvent an issue where the compiler
            // allocates stack space for every case branch when no optimizations are
            // enabled. https://github.com/apple/swift-protobuf/issues/1034
            switch fieldNumber {
            case 1: try { try decoder.decodeSingularStringField(value: &self.name) }()
            case 2: try { try decoder.decodeSingularInt64Field(value: &self.id) }()
            case 3: try { try decoder.decodeSingularFloatField(value: &self.price) }()
            case 4: try { try decoder.decodeSingularInt64Field(value: &self.categoryID) }()
            case 5: try { try decoder.decodeSingularInt64Field(value: &self.orderItemID) }()
            case 6: try { try decoder.decodeRepeatedMessageField(value: &self.options) }()
            default: break
            }
        }
    }

    public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
        if !name.isEmpty {
            try visitor.visitSingularStringField(value: name, fieldNumber: 1)
        }
        if id != 0 {
            try visitor.visitSingularInt64Field(value: id, fieldNumber: 2)
        }
        if price != 0 {
            try visitor.visitSingularFloatField(value: price, fieldNumber: 3)
        }
        if categoryID != 0 {
            try visitor.visitSingularInt64Field(value: categoryID, fieldNumber: 4)
        }
        if orderItemID != 0 {
            try visitor.visitSingularInt64Field(value: orderItemID, fieldNumber: 5)
        }
        if !options.isEmpty {
            try visitor.visitRepeatedMessageField(value: options, fieldNumber: 6)
        }
        try unknownFields.traverse(visitor: &visitor)
    }

    public static func == (lhs: Tospb_Item, rhs: Tospb_Item) -> Bool {
        if lhs.name != rhs.name { return false }
        if lhs.id != rhs.id { return false }
        if lhs.price != rhs.price { return false }
        if lhs.categoryID != rhs.categoryID { return false }
        if lhs.orderItemID != rhs.orderItemID { return false }
        if lhs.options != rhs.options { return false }
        if lhs.unknownFields != rhs.unknownFields { return false }
        return true
    }
}

extension Tospb_Option: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
    public static let protoMessageName: String = _protobuf_package + ".Option"
    public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
        1: .same(proto: "name"),
        2: .same(proto: "id"),
        3: .same(proto: "price"),
        4: .same(proto: "selected"),
    ]

    public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
        while let fieldNumber = try decoder.nextFieldNumber() {
            // The use of inline closures is to circumvent an issue where the compiler
            // allocates stack space for every case branch when no optimizations are
            // enabled. https://github.com/apple/swift-protobuf/issues/1034
            switch fieldNumber {
            case 1: try { try decoder.decodeSingularStringField(value: &self.name) }()
            case 2: try { try decoder.decodeSingularInt64Field(value: &self.id) }()
            case 3: try { try decoder.decodeSingularFloatField(value: &self.price) }()
            case 4: try { try decoder.decodeSingularBoolField(value: &self.selected) }()
            default: break
            }
        }
    }

    public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
        if !name.isEmpty {
            try visitor.visitSingularStringField(value: name, fieldNumber: 1)
        }
        if id != 0 {
            try visitor.visitSingularInt64Field(value: id, fieldNumber: 2)
        }
        if price != 0 {
            try visitor.visitSingularFloatField(value: price, fieldNumber: 3)
        }
        if selected != false {
            try visitor.visitSingularBoolField(value: selected, fieldNumber: 4)
        }
        try unknownFields.traverse(visitor: &visitor)
    }

    public static func == (lhs: Tospb_Option, rhs: Tospb_Option) -> Bool {
        if lhs.name != rhs.name { return false }
        if lhs.id != rhs.id { return false }
        if lhs.price != rhs.price { return false }
        if lhs.selected != rhs.selected { return false }
        if lhs.unknownFields != rhs.unknownFields { return false }
        return true
    }
}

extension Tospb_Category: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
    public static let protoMessageName: String = _protobuf_package + ".Category"
    public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
        1: .same(proto: "name"),
        2: .same(proto: "id"),
        3: .same(proto: "items"),
    ]

    public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
        while let fieldNumber = try decoder.nextFieldNumber() {
            // The use of inline closures is to circumvent an issue where the compiler
            // allocates stack space for every case branch when no optimizations are
            // enabled. https://github.com/apple/swift-protobuf/issues/1034
            switch fieldNumber {
            case 1: try { try decoder.decodeSingularStringField(value: &self.name) }()
            case 2: try { try decoder.decodeSingularInt64Field(value: &self.id) }()
            case 3: try { try decoder.decodeRepeatedMessageField(value: &self.items) }()
            default: break
            }
        }
    }

    public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
        if !name.isEmpty {
            try visitor.visitSingularStringField(value: name, fieldNumber: 1)
        }
        if id != 0 {
            try visitor.visitSingularInt64Field(value: id, fieldNumber: 2)
        }
        if !items.isEmpty {
            try visitor.visitRepeatedMessageField(value: items, fieldNumber: 3)
        }
        try unknownFields.traverse(visitor: &visitor)
    }

    public static func == (lhs: Tospb_Category, rhs: Tospb_Category) -> Bool {
        if lhs.name != rhs.name { return false }
        if lhs.id != rhs.id { return false }
        if lhs.items != rhs.items { return false }
        if lhs.unknownFields != rhs.unknownFields { return false }
        return true
    }
}

extension Tospb_Menu: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
    public static let protoMessageName: String = _protobuf_package + ".Menu"
    public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
        1: .same(proto: "categories"),
    ]

    public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
        while let fieldNumber = try decoder.nextFieldNumber() {
            // The use of inline closures is to circumvent an issue where the compiler
            // allocates stack space for every case branch when no optimizations are
            // enabled. https://github.com/apple/swift-protobuf/issues/1034
            switch fieldNumber {
            case 1: try { try decoder.decodeRepeatedMessageField(value: &self.categories) }()
            default: break
            }
        }
    }

    public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
        if !categories.isEmpty {
            try visitor.visitRepeatedMessageField(value: categories, fieldNumber: 1)
        }
        try unknownFields.traverse(visitor: &visitor)
    }

    public static func == (lhs: Tospb_Menu, rhs: Tospb_Menu) -> Bool {
        if lhs.categories != rhs.categories { return false }
        if lhs.unknownFields != rhs.unknownFields { return false }
        return true
    }
}

extension Tospb_Order: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
    public static let protoMessageName: String = _protobuf_package + ".Order"
    public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
        1: .same(proto: "id"),
        2: .same(proto: "name"),
        3: .same(proto: "items"),
        4: .same(proto: "total"),
        5: .same(proto: "status"),
        6: .standard(proto: "time_ordered"),
        7: .standard(proto: "time_complete"),
    ]

    public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
        while let fieldNumber = try decoder.nextFieldNumber() {
            // The use of inline closures is to circumvent an issue where the compiler
            // allocates stack space for every case branch when no optimizations are
            // enabled. https://github.com/apple/swift-protobuf/issues/1034
            switch fieldNumber {
            case 1: try { try decoder.decodeSingularInt64Field(value: &self.id) }()
            case 2: try { try decoder.decodeSingularStringField(value: &self.name) }()
            case 3: try { try decoder.decodeRepeatedMessageField(value: &self.items) }()
            case 4: try { try decoder.decodeSingularFloatField(value: &self.total) }()
            case 5: try { try decoder.decodeSingularStringField(value: &self.status) }()
            case 6: try { try decoder.decodeSingularStringField(value: &self.timeOrdered) }()
            case 7: try { try decoder.decodeSingularStringField(value: &self.timeComplete) }()
            default: break
            }
        }
    }

    public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
        if id != 0 {
            try visitor.visitSingularInt64Field(value: id, fieldNumber: 1)
        }
        if !name.isEmpty {
            try visitor.visitSingularStringField(value: name, fieldNumber: 2)
        }
        if !items.isEmpty {
            try visitor.visitRepeatedMessageField(value: items, fieldNumber: 3)
        }
        if total != 0 {
            try visitor.visitSingularFloatField(value: total, fieldNumber: 4)
        }
        if !status.isEmpty {
            try visitor.visitSingularStringField(value: status, fieldNumber: 5)
        }
        if !timeOrdered.isEmpty {
            try visitor.visitSingularStringField(value: timeOrdered, fieldNumber: 6)
        }
        if !timeComplete.isEmpty {
            try visitor.visitSingularStringField(value: timeComplete, fieldNumber: 7)
        }
        try unknownFields.traverse(visitor: &visitor)
    }

    public static func == (lhs: Tospb_Order, rhs: Tospb_Order) -> Bool {
        if lhs.id != rhs.id { return false }
        if lhs.name != rhs.name { return false }
        if lhs.items != rhs.items { return false }
        if lhs.total != rhs.total { return false }
        if lhs.status != rhs.status { return false }
        if lhs.timeOrdered != rhs.timeOrdered { return false }
        if lhs.timeComplete != rhs.timeComplete { return false }
        if lhs.unknownFields != rhs.unknownFields { return false }
        return true
    }
}

extension Tospb_Response: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
    public static let protoMessageName: String = _protobuf_package + ".Response"
    public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
        1: .same(proto: "response"),
    ]

    public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
        while let fieldNumber = try decoder.nextFieldNumber() {
            // The use of inline closures is to circumvent an issue where the compiler
            // allocates stack space for every case branch when no optimizations are
            // enabled. https://github.com/apple/swift-protobuf/issues/1034
            switch fieldNumber {
            case 1: try { try decoder.decodeSingularStringField(value: &self.response) }()
            default: break
            }
        }
    }

    public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
        if !response.isEmpty {
            try visitor.visitSingularStringField(value: response, fieldNumber: 1)
        }
        try unknownFields.traverse(visitor: &visitor)
    }

    public static func == (lhs: Tospb_Response, rhs: Tospb_Response) -> Bool {
        if lhs.response != rhs.response { return false }
        if lhs.unknownFields != rhs.unknownFields { return false }
        return true
    }
}

extension Tospb_CreateMenuItemResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
    public static let protoMessageName: String = _protobuf_package + ".CreateMenuItemResponse"
    public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
        1: .same(proto: "id"),
    ]

    public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
        while let fieldNumber = try decoder.nextFieldNumber() {
            // The use of inline closures is to circumvent an issue where the compiler
            // allocates stack space for every case branch when no optimizations are
            // enabled. https://github.com/apple/swift-protobuf/issues/1034
            switch fieldNumber {
            case 1: try { try decoder.decodeSingularInt64Field(value: &self.id) }()
            default: break
            }
        }
    }

    public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
        if id != 0 {
            try visitor.visitSingularInt64Field(value: id, fieldNumber: 1)
        }
        try unknownFields.traverse(visitor: &visitor)
    }

    public static func == (lhs: Tospb_CreateMenuItemResponse, rhs: Tospb_CreateMenuItemResponse) -> Bool {
        if lhs.id != rhs.id { return false }
        if lhs.unknownFields != rhs.unknownFields { return false }
        return true
    }
}

extension Tospb_CompleteOrderRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
    public static let protoMessageName: String = _protobuf_package + ".CompleteOrderRequest"
    public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
        1: .same(proto: "id"),
    ]

    public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
        while let fieldNumber = try decoder.nextFieldNumber() {
            // The use of inline closures is to circumvent an issue where the compiler
            // allocates stack space for every case branch when no optimizations are
            // enabled. https://github.com/apple/swift-protobuf/issues/1034
            switch fieldNumber {
            case 1: try { try decoder.decodeSingularInt64Field(value: &self.id) }()
            default: break
            }
        }
    }

    public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
        if id != 0 {
            try visitor.visitSingularInt64Field(value: id, fieldNumber: 1)
        }
        try unknownFields.traverse(visitor: &visitor)
    }

    public static func == (lhs: Tospb_CompleteOrderRequest, rhs: Tospb_CompleteOrderRequest) -> Bool {
        if lhs.id != rhs.id { return false }
        if lhs.unknownFields != rhs.unknownFields { return false }
        return true
    }
}

extension Tospb_OrdersRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
    public static let protoMessageName: String = _protobuf_package + ".OrdersRequest"
    public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
        1: .same(proto: "request"),
    ]

    public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
        while let fieldNumber = try decoder.nextFieldNumber() {
            // The use of inline closures is to circumvent an issue where the compiler
            // allocates stack space for every case branch when no optimizations are
            // enabled. https://github.com/apple/swift-protobuf/issues/1034
            switch fieldNumber {
            case 1: try { try decoder.decodeSingularStringField(value: &self.request) }()
            default: break
            }
        }
    }

    public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
        if !request.isEmpty {
            try visitor.visitSingularStringField(value: request, fieldNumber: 1)
        }
        try unknownFields.traverse(visitor: &visitor)
    }

    public static func == (lhs: Tospb_OrdersRequest, rhs: Tospb_OrdersRequest) -> Bool {
        if lhs.request != rhs.request { return false }
        if lhs.unknownFields != rhs.unknownFields { return false }
        return true
    }
}

extension Tospb_OrdersResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
    public static let protoMessageName: String = _protobuf_package + ".OrdersResponse"
    public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
        1: .same(proto: "orders"),
    ]

    public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
        while let fieldNumber = try decoder.nextFieldNumber() {
            // The use of inline closures is to circumvent an issue where the compiler
            // allocates stack space for every case branch when no optimizations are
            // enabled. https://github.com/apple/swift-protobuf/issues/1034
            switch fieldNumber {
            case 1: try { try decoder.decodeRepeatedMessageField(value: &self.orders) }()
            default: break
            }
        }
    }

    public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
        if !orders.isEmpty {
            try visitor.visitRepeatedMessageField(value: orders, fieldNumber: 1)
        }
        try unknownFields.traverse(visitor: &visitor)
    }

    public static func == (lhs: Tospb_OrdersResponse, rhs: Tospb_OrdersResponse) -> Bool {
        if lhs.orders != rhs.orders { return false }
        if lhs.unknownFields != rhs.unknownFields { return false }
        return true
    }
}

extension Tospb_Empty: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
    public static let protoMessageName: String = _protobuf_package + ".Empty"
    public static let _protobuf_nameMap = SwiftProtobuf._NameMap()

    public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
        while let _ = try decoder.nextFieldNumber() {}
    }

    public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
        try unknownFields.traverse(visitor: &visitor)
    }

    public static func == (lhs: Tospb_Empty, rhs: Tospb_Empty) -> Bool {
        if lhs.unknownFields != rhs.unknownFields { return false }
        return true
    }
}

extension Tospb_DeleteMenuItemRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
    public static let protoMessageName: String = _protobuf_package + ".DeleteMenuItemRequest"
    public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
        1: .same(proto: "id"),
    ]

    public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
        while let fieldNumber = try decoder.nextFieldNumber() {
            // The use of inline closures is to circumvent an issue where the compiler
            // allocates stack space for every case branch when no optimizations are
            // enabled. https://github.com/apple/swift-protobuf/issues/1034
            switch fieldNumber {
            case 1: try { try decoder.decodeSingularInt64Field(value: &self.id) }()
            default: break
            }
        }
    }

    public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
        if id != 0 {
            try visitor.visitSingularInt64Field(value: id, fieldNumber: 1)
        }
        try unknownFields.traverse(visitor: &visitor)
    }

    public static func == (lhs: Tospb_DeleteMenuItemRequest, rhs: Tospb_DeleteMenuItemRequest) -> Bool {
        if lhs.id != rhs.id { return false }
        if lhs.unknownFields != rhs.unknownFields { return false }
        return true
    }
}

extension Tospb_Ping: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
    public static let protoMessageName: String = _protobuf_package + ".Ping"
    public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
        1: .same(proto: "message"),
        2: .same(proto: "delaySeconds"),
    ]

    public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
        while let fieldNumber = try decoder.nextFieldNumber() {
            // The use of inline closures is to circumvent an issue where the compiler
            // allocates stack space for every case branch when no optimizations are
            // enabled. https://github.com/apple/swift-protobuf/issues/1034
            switch fieldNumber {
            case 1: try { try decoder.decodeSingularStringField(value: &self.message) }()
            case 2: try { try decoder.decodeSingularFloatField(value: &self.delaySeconds) }()
            default: break
            }
        }
    }

    public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
        if !message.isEmpty {
            try visitor.visitSingularStringField(value: message, fieldNumber: 1)
        }
        if delaySeconds != 0 {
            try visitor.visitSingularFloatField(value: delaySeconds, fieldNumber: 2)
        }
        try unknownFields.traverse(visitor: &visitor)
    }

    public static func == (lhs: Tospb_Ping, rhs: Tospb_Ping) -> Bool {
        if lhs.message != rhs.message { return false }
        if lhs.delaySeconds != rhs.delaySeconds { return false }
        if lhs.unknownFields != rhs.unknownFields { return false }
        return true
    }
}

extension Tospb_Pong: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
    public static let protoMessageName: String = _protobuf_package + ".Pong"
    public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
        1: .same(proto: "message"),
    ]

    public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
        while let fieldNumber = try decoder.nextFieldNumber() {
            // The use of inline closures is to circumvent an issue where the compiler
            // allocates stack space for every case branch when no optimizations are
            // enabled. https://github.com/apple/swift-protobuf/issues/1034
            switch fieldNumber {
            case 1: try { try decoder.decodeSingularStringField(value: &self.message) }()
            default: break
            }
        }
    }

    public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
        if !message.isEmpty {
            try visitor.visitSingularStringField(value: message, fieldNumber: 1)
        }
        try unknownFields.traverse(visitor: &visitor)
    }

    public static func == (lhs: Tospb_Pong, rhs: Tospb_Pong) -> Bool {
        if lhs.message != rhs.message { return false }
        if lhs.unknownFields != rhs.unknownFields { return false }
        return true
    }
}
