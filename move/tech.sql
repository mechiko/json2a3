SET NOCOUNT ON;
BEGIN TRY
  BEGIN TRANSACTION;
 delete from product_guides;
 delete from order_mark_codes;
 delete from order_mark_codes_serial_numbers ;
 delete from order_mark_utilisation ;
 delete from order_mark_utilisation_codes ;
 delete from order_mark_commissioning;
 delete from order_mark_commissioning_aggregations ;
 delete from order_mark_commissioning_atk ;
 delete from order_mark_commissioning_codes ;
 delete from order_mark_aggregation;
 delete from order_mark_aggregation_codes;
 delete from order_mark_atk ;
 delete from order_mark_atk_codes;
 delete from order_mark_connecting_tap ;
 delete from order_mark_connecting_tap_codes ;
 delete from order_mark_service_providers;
 
IF EXISTS (
  SELECT 1 FROM sys.identity_columns 
  WHERE object_id = OBJECT_ID(N'[dbo].[product_guides]')
)
  DBCC CHECKIDENT (N'[dbo].[product_guides]', RESEED, 0);

IF EXISTS (
  SELECT 1 FROM sys.identity_columns 
  WHERE object_id = OBJECT_ID(N'[dbo].[order_mark_codes]')
)
 DBCC CHECKIDENT ([order_mark_codes], RESEED, 0);

IF EXISTS (
  SELECT 1 FROM sys.identity_columns 
  WHERE object_id = OBJECT_ID(N'[dbo].[order_mark_codes_serial_numbers]')
)
 DBCC CHECKIDENT ([order_mark_codes_serial_numbers], RESEED, 0);

IF EXISTS (
  SELECT 1 FROM sys.identity_columns 
  WHERE object_id = OBJECT_ID(N'[dbo].[order_mark_utilisation]')
)
 DBCC CHECKIDENT ([order_mark_utilisation], RESEED, 0);

IF EXISTS (
  SELECT 1 FROM sys.identity_columns 
  WHERE object_id = OBJECT_ID(N'[dbo].[order_mark_utilisation_codes]')
)
 DBCC CHECKIDENT ([order_mark_utilisation_codes], RESEED, 0);

IF EXISTS (
  SELECT 1 FROM sys.identity_columns 
  WHERE object_id = OBJECT_ID(N'[dbo].[order_mark_commissioning]')
)
 DBCC CHECKIDENT ([order_mark_commissioning], RESEED, 0);

IF EXISTS (
  SELECT 1 FROM sys.identity_columns 
  WHERE object_id = OBJECT_ID(N'[dbo].[order_mark_commissioning_aggregations]')
)
 DBCC CHECKIDENT ([order_mark_commissioning_aggregations], RESEED, 0);

IF EXISTS (
  SELECT 1 FROM sys.identity_columns 
  WHERE object_id = OBJECT_ID(N'[dbo].[order_mark_commissioning_atk]')
)
 DBCC CHECKIDENT ([order_mark_commissioning_atk], RESEED, 0);

IF EXISTS (
  SELECT 1 FROM sys.identity_columns 
  WHERE object_id = OBJECT_ID(N'[dbo].[order_mark_commissioning_codes]')
)
 DBCC CHECKIDENT ([order_mark_commissioning_codes], RESEED, 0);

IF EXISTS (
  SELECT 1 FROM sys.identity_columns 
  WHERE object_id = OBJECT_ID(N'[dbo].[order_mark_aggregation]')
)
 DBCC CHECKIDENT ([order_mark_aggregation], RESEED, 0);

IF EXISTS (
  SELECT 1 FROM sys.identity_columns 
  WHERE object_id = OBJECT_ID(N'[dbo].[order_mark_aggregation_codes]')
)
 DBCC CHECKIDENT ([order_mark_aggregation_codes], RESEED, 0);

IF EXISTS (
  SELECT 1 FROM sys.identity_columns 
  WHERE object_id = OBJECT_ID(N'[dbo].[order_mark_atk]')
)
 DBCC CHECKIDENT ([order_mark_atk], RESEED, 0);

IF EXISTS (
  SELECT 1 FROM sys.identity_columns 
  WHERE object_id = OBJECT_ID(N'[dbo].[order_mark_atk_codes]')
)
 DBCC CHECKIDENT ([order_mark_atk_codes], RESEED, 0);

IF EXISTS (
  SELECT 1 FROM sys.identity_columns 
  WHERE object_id = OBJECT_ID(N'[dbo].[order_mark_connecting_tap]')
)
 DBCC CHECKIDENT ([order_mark_connecting_tap], RESEED, 0);

IF EXISTS (
  SELECT 1 FROM sys.identity_columns 
  WHERE object_id = OBJECT_ID(N'[dbo].[order_mark_connecting_tap_codes]')
)
 DBCC CHECKIDENT ([order_mark_connecting_tap_codes], RESEED, 0);

IF EXISTS (
  SELECT 1 FROM sys.identity_columns 
  WHERE object_id = OBJECT_ID(N'[dbo].[order_mark_service_providers]')
)
 DBCC CHECKIDENT ([order_mark_service_providers], RESEED, 0);
  COMMIT TRANSACTION;
END TRY
BEGIN CATCH
  IF @@TRANCOUNT > 0 ROLLBACK TRANSACTION;
  THROW;
END CATCH
