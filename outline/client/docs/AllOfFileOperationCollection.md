# AllOfFileOperationCollection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the object. | [optional] [default to null]
**Name** | **string** | The name of the collection. | [optional] [default to null]
**Description** | **string** | A description of the collection, may contain markdown formatting | [optional] [default to null]
**Sort** | [***CollectionSort**](Collection_sort.md) |  | [optional] [default to null]
**Index** | **string** | The position of the collection in the sidebar | [optional] [default to null]
**Documents** | [**[]NavigationNode**](NavigationNode.md) |  | [optional] [default to null]
**Color** | **string** | A color representing the collection, this is used to help make collections more identifiable in the UI. It should be in HEX format including the # | [optional] [default to null]
**Icon** | **string** | A string that represents an icon in the outline-icons package | [optional] [default to null]
**Permission** | **string** |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time that this object was created | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time that this object was last changed | [optional] [default to null]
**DeletedAt** | [**time.Time**](time.Time.md) | The date and time that this object was deleted | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

