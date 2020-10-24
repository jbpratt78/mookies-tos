//
// DO NOT EDIT.
//
// Generated by the protocol buffer compiler.
// Source: protofiles/tos.proto
//

//
// Copyright 2018, gRPC Authors All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
import GRPC
import NIO
import SwiftProtobuf

/// Usage: instantiate Tospb_MenuServiceClient, then call methods of this protocol to make API calls.
public protocol Tospb_MenuServiceClientProtocol: GRPCClient {
    func getMenu(
        _ request: Tospb_Empty,
        callOptions: CallOptions?
    ) -> UnaryCall<Tospb_Empty, Tospb_Menu>

    func createMenuItem(
        _ request: Tospb_Item,
        callOptions: CallOptions?
    ) -> UnaryCall<Tospb_Item, Tospb_CreateMenuItemResponse>

    func updateMenuItem(
        _ request: Tospb_Item,
        callOptions: CallOptions?
    ) -> UnaryCall<Tospb_Item, Tospb_Response>

    func deleteMenuItem(
        _ request: Tospb_DeleteMenuItemRequest,
        callOptions: CallOptions?
    ) -> UnaryCall<Tospb_DeleteMenuItemRequest, Tospb_Response>

    func createMenuItemOption(
        _ request: Tospb_Option,
        callOptions: CallOptions?
    ) -> UnaryCall<Tospb_Option, Tospb_Response>
}

extension Tospb_MenuServiceClientProtocol {
    /// Unary call to GetMenu
    ///
    /// - Parameters:
    ///   - request: Request to send to GetMenu.
    ///   - callOptions: Call options.
    /// - Returns: A `UnaryCall` with futures for the metadata, status and response.
    public func getMenu(
        _ request: Tospb_Empty,
        callOptions: CallOptions? = nil
    ) -> UnaryCall<Tospb_Empty, Tospb_Menu> {
        return makeUnaryCall(
            path: "/tospb.MenuService/GetMenu",
            request: request,
            callOptions: callOptions ?? defaultCallOptions
        )
    }

    /// Unary call to CreateMenuItem
    ///
    /// - Parameters:
    ///   - request: Request to send to CreateMenuItem.
    ///   - callOptions: Call options.
    /// - Returns: A `UnaryCall` with futures for the metadata, status and response.
    public func createMenuItem(
        _ request: Tospb_Item,
        callOptions: CallOptions? = nil
    ) -> UnaryCall<Tospb_Item, Tospb_CreateMenuItemResponse> {
        return makeUnaryCall(
            path: "/tospb.MenuService/CreateMenuItem",
            request: request,
            callOptions: callOptions ?? defaultCallOptions
        )
    }

    /// Unary call to UpdateMenuItem
    ///
    /// - Parameters:
    ///   - request: Request to send to UpdateMenuItem.
    ///   - callOptions: Call options.
    /// - Returns: A `UnaryCall` with futures for the metadata, status and response.
    public func updateMenuItem(
        _ request: Tospb_Item,
        callOptions: CallOptions? = nil
    ) -> UnaryCall<Tospb_Item, Tospb_Response> {
        return makeUnaryCall(
            path: "/tospb.MenuService/UpdateMenuItem",
            request: request,
            callOptions: callOptions ?? defaultCallOptions
        )
    }

    /// Unary call to DeleteMenuItem
    ///
    /// - Parameters:
    ///   - request: Request to send to DeleteMenuItem.
    ///   - callOptions: Call options.
    /// - Returns: A `UnaryCall` with futures for the metadata, status and response.
    public func deleteMenuItem(
        _ request: Tospb_DeleteMenuItemRequest,
        callOptions: CallOptions? = nil
    ) -> UnaryCall<Tospb_DeleteMenuItemRequest, Tospb_Response> {
        return makeUnaryCall(
            path: "/tospb.MenuService/DeleteMenuItem",
            request: request,
            callOptions: callOptions ?? defaultCallOptions
        )
    }

    /// Unary call to CreateMenuItemOption
    ///
    /// - Parameters:
    ///   - request: Request to send to CreateMenuItemOption.
    ///   - callOptions: Call options.
    /// - Returns: A `UnaryCall` with futures for the metadata, status and response.
    public func createMenuItemOption(
        _ request: Tospb_Option,
        callOptions: CallOptions? = nil
    ) -> UnaryCall<Tospb_Option, Tospb_Response> {
        return makeUnaryCall(
            path: "/tospb.MenuService/CreateMenuItemOption",
            request: request,
            callOptions: callOptions ?? defaultCallOptions
        )
    }
}

public final class Tospb_MenuServiceClient: Tospb_MenuServiceClientProtocol {
    public let channel: GRPCChannel
    public var defaultCallOptions: CallOptions

    /// Creates a client for the tospb.MenuService service.
    ///
    /// - Parameters:
    ///   - channel: `GRPCChannel` to the service host.
    ///   - defaultCallOptions: Options to use for each service call if the user doesn't provide them.
    public init(channel: GRPCChannel, defaultCallOptions: CallOptions = CallOptions()) {
        self.channel = channel
        self.defaultCallOptions = defaultCallOptions
    }
}

/// Usage: instantiate Tospb_OrderServiceClient, then call methods of this protocol to make API calls.
public protocol Tospb_OrderServiceClientProtocol: GRPCClient {
    func submitOrder(
        _ request: Tospb_Order,
        callOptions: CallOptions?
    ) -> UnaryCall<Tospb_Order, Tospb_Response>

    func activeOrders(
        _ request: Tospb_Empty,
        callOptions: CallOptions?
    ) -> UnaryCall<Tospb_Empty, Tospb_OrdersResponse>

    func completeOrder(
        _ request: Tospb_CompleteOrderRequest,
        callOptions: CallOptions?
    ) -> UnaryCall<Tospb_CompleteOrderRequest, Tospb_Response>

    func subscribeToOrders(
        _ request: Tospb_Empty,
        callOptions: CallOptions?,
        handler: @escaping (Tospb_Order) -> Void
    ) -> ServerStreamingCall<Tospb_Empty, Tospb_Order>
}

extension Tospb_OrderServiceClientProtocol {
    /// Unary
    ///
    /// - Parameters:
    ///   - request: Request to send to SubmitOrder.
    ///   - callOptions: Call options.
    /// - Returns: A `UnaryCall` with futures for the metadata, status and response.
    public func submitOrder(
        _ request: Tospb_Order,
        callOptions: CallOptions? = nil
    ) -> UnaryCall<Tospb_Order, Tospb_Response> {
        return makeUnaryCall(
            path: "/tospb.OrderService/SubmitOrder",
            request: request,
            callOptions: callOptions ?? defaultCallOptions
        )
    }

    /// Unary call to ActiveOrders
    ///
    /// - Parameters:
    ///   - request: Request to send to ActiveOrders.
    ///   - callOptions: Call options.
    /// - Returns: A `UnaryCall` with futures for the metadata, status and response.
    public func activeOrders(
        _ request: Tospb_Empty,
        callOptions: CallOptions? = nil
    ) -> UnaryCall<Tospb_Empty, Tospb_OrdersResponse> {
        return makeUnaryCall(
            path: "/tospb.OrderService/ActiveOrders",
            request: request,
            callOptions: callOptions ?? defaultCallOptions
        )
    }

    /// Unary call to CompleteOrder
    ///
    /// - Parameters:
    ///   - request: Request to send to CompleteOrder.
    ///   - callOptions: Call options.
    /// - Returns: A `UnaryCall` with futures for the metadata, status and response.
    public func completeOrder(
        _ request: Tospb_CompleteOrderRequest,
        callOptions: CallOptions? = nil
    ) -> UnaryCall<Tospb_CompleteOrderRequest, Tospb_Response> {
        return makeUnaryCall(
            path: "/tospb.OrderService/CompleteOrder",
            request: request,
            callOptions: callOptions ?? defaultCallOptions
        )
    }

    /// server streaming
    ///
    /// - Parameters:
    ///   - request: Request to send to SubscribeToOrders.
    ///   - callOptions: Call options.
    ///   - handler: A closure called when each response is received from the server.
    /// - Returns: A `ServerStreamingCall` with futures for the metadata and status.
    public func subscribeToOrders(
        _ request: Tospb_Empty,
        callOptions: CallOptions? = nil,
        handler: @escaping (Tospb_Order) -> Void
    ) -> ServerStreamingCall<Tospb_Empty, Tospb_Order> {
        return makeServerStreamingCall(
            path: "/tospb.OrderService/SubscribeToOrders",
            request: request,
            callOptions: callOptions ?? defaultCallOptions,
            handler: handler
        )
    }
}

public final class Tospb_OrderServiceClient: Tospb_OrderServiceClientProtocol {
    public let channel: GRPCChannel
    public var defaultCallOptions: CallOptions

    /// Creates a client for the tospb.OrderService service.
    ///
    /// - Parameters:
    ///   - channel: `GRPCChannel` to the service host.
    ///   - defaultCallOptions: Options to use for each service call if the user doesn't provide them.
    public init(channel: GRPCChannel, defaultCallOptions: CallOptions = CallOptions()) {
        self.channel = channel
        self.defaultCallOptions = defaultCallOptions
    }
}

/// To build a server, implement a class that conforms to this protocol.
public protocol Tospb_MenuServiceProvider: CallHandlerProvider {
    func getMenu(request: Tospb_Empty, context: StatusOnlyCallContext) -> EventLoopFuture<Tospb_Menu>
    func createMenuItem(request: Tospb_Item, context: StatusOnlyCallContext) -> EventLoopFuture<Tospb_CreateMenuItemResponse>
    func updateMenuItem(request: Tospb_Item, context: StatusOnlyCallContext) -> EventLoopFuture<Tospb_Response>
    func deleteMenuItem(request: Tospb_DeleteMenuItemRequest, context: StatusOnlyCallContext) -> EventLoopFuture<Tospb_Response>
    func createMenuItemOption(request: Tospb_Option, context: StatusOnlyCallContext) -> EventLoopFuture<Tospb_Response>
}

extension Tospb_MenuServiceProvider {
    public var serviceName: Substring { return "tospb.MenuService" }

    /// Determines, calls and returns the appropriate request handler, depending on the request's method.
    /// Returns nil for methods not handled by this service.
    public func handleMethod(_ methodName: Substring, callHandlerContext: CallHandlerContext) -> GRPCCallHandler? {
        switch methodName {
        case "GetMenu":
            return CallHandlerFactory.makeUnary(callHandlerContext: callHandlerContext) { context in
                { request in
                    self.getMenu(request: request, context: context)
                }
            }

        case "CreateMenuItem":
            return CallHandlerFactory.makeUnary(callHandlerContext: callHandlerContext) { context in
                { request in
                    self.createMenuItem(request: request, context: context)
                }
            }

        case "UpdateMenuItem":
            return CallHandlerFactory.makeUnary(callHandlerContext: callHandlerContext) { context in
                { request in
                    self.updateMenuItem(request: request, context: context)
                }
            }

        case "DeleteMenuItem":
            return CallHandlerFactory.makeUnary(callHandlerContext: callHandlerContext) { context in
                { request in
                    self.deleteMenuItem(request: request, context: context)
                }
            }

        case "CreateMenuItemOption":
            return CallHandlerFactory.makeUnary(callHandlerContext: callHandlerContext) { context in
                { request in
                    self.createMenuItemOption(request: request, context: context)
                }
            }

        default: return nil
        }
    }
}

/// To build a server, implement a class that conforms to this protocol.
public protocol Tospb_OrderServiceProvider: CallHandlerProvider {
    /// Unary
    func submitOrder(request: Tospb_Order, context: StatusOnlyCallContext) -> EventLoopFuture<Tospb_Response>
    func activeOrders(request: Tospb_Empty, context: StatusOnlyCallContext) -> EventLoopFuture<Tospb_OrdersResponse>
    func completeOrder(request: Tospb_CompleteOrderRequest, context: StatusOnlyCallContext) -> EventLoopFuture<Tospb_Response>
    /// server streaming
    func subscribeToOrders(request: Tospb_Empty, context: StreamingResponseCallContext<Tospb_Order>) -> EventLoopFuture<GRPCStatus>
}

extension Tospb_OrderServiceProvider {
    public var serviceName: Substring { return "tospb.OrderService" }

    /// Determines, calls and returns the appropriate request handler, depending on the request's method.
    /// Returns nil for methods not handled by this service.
    public func handleMethod(_ methodName: Substring, callHandlerContext: CallHandlerContext) -> GRPCCallHandler? {
        switch methodName {
        case "SubmitOrder":
            return CallHandlerFactory.makeUnary(callHandlerContext: callHandlerContext) { context in
                { request in
                    self.submitOrder(request: request, context: context)
                }
            }

        case "ActiveOrders":
            return CallHandlerFactory.makeUnary(callHandlerContext: callHandlerContext) { context in
                { request in
                    self.activeOrders(request: request, context: context)
                }
            }

        case "CompleteOrder":
            return CallHandlerFactory.makeUnary(callHandlerContext: callHandlerContext) { context in
                { request in
                    self.completeOrder(request: request, context: context)
                }
            }

        case "SubscribeToOrders":
            return CallHandlerFactory.makeServerStreaming(callHandlerContext: callHandlerContext) { context in
                { request in
                    self.subscribeToOrders(request: request, context: context)
                }
            }

        default: return nil
        }
    }
}