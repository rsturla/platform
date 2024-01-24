// my-service/test.proto

// @generated by protoc-gen-es v1.6.0
// @generated from file test/v1/test.proto (package test.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message test.v1.Person
 */
export declare class Person extends Message<Person> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: int32 id = 2;
   */
  id: number;

  /**
   * @generated from field: string email = 3;
   */
  email: string;

  /**
   * @generated from field: repeated test.v1.Person.PhoneNumber phones = 4;
   */
  phones: Person_PhoneNumber[];

  constructor(data?: PartialMessage<Person>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "test.v1.Person";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Person;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Person;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Person;

  static equals(a: Person | PlainMessage<Person> | undefined, b: Person | PlainMessage<Person> | undefined): boolean;
}

/**
 * @generated from enum test.v1.Person.PhoneType
 */
export declare enum Person_PhoneType {
  /**
   * @generated from enum value: PHONE_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: PHONE_TYPE_MOBILE = 1;
   */
  MOBILE = 1,

  /**
   * @generated from enum value: PHONE_TYPE_HOME = 2;
   */
  HOME = 2,

  /**
   * @generated from enum value: PHONE_TYPE_WORK = 3;
   */
  WORK = 3,
}

/**
 * @generated from message test.v1.Person.PhoneNumber
 */
export declare class Person_PhoneNumber extends Message<Person_PhoneNumber> {
  /**
   * @generated from field: string number = 1;
   */
  number: string;

  /**
   * @generated from field: test.v1.Person.PhoneType type = 2;
   */
  type: Person_PhoneType;

  constructor(data?: PartialMessage<Person_PhoneNumber>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "test.v1.Person.PhoneNumber";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Person_PhoneNumber;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Person_PhoneNumber;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Person_PhoneNumber;

  static equals(a: Person_PhoneNumber | PlainMessage<Person_PhoneNumber> | undefined, b: Person_PhoneNumber | PlainMessage<Person_PhoneNumber> | undefined): boolean;
}
