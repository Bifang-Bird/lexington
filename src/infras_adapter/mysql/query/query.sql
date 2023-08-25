-- name: GetAll :many
select
    o.id,
    o.order_id,
    o.user_id,
    o.payment_amount,
    o.payment_cycle,
    o.payment_create_time,
    o.payment_end_time,
    o.payment_order_status,
    o.payment_cancel_details,
    o.payment_context,
    o.currency
FROM
    payment.payment_order o;

-- name: GetByID :one

select
    o.id,
    o.order_id,
    o.user_id,
    o.payment_amount,
    o.payment_cycle,
    o.payment_create_time,
    o.payment_end_time,
    o.payment_order_status,
    o.payment_cancel_details,
    o.payment_context,
    o.currency
FROM
    payment.payment_order o
WHERE o.id = ? LIMIT 1;


-- name: CreatePaymentOrder :execresult
INSERT INTO
    payment.payment_order (
             id,
             order_id,
             user_id,
             payment_amount,
             payment_cycle,
             payment_create_time,
             payment_end_time,
             payment_order_status,
             payment_cancel_details,
             payment_context,
             currency
             )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?);

-- name: SyncPaymentFlow :exec
UPDATE payment.payment_flow
SET
    payment_flow_status = ?,
    payment_sync_deal_result =?,
    payment_fail_reason = ?,
    update_time = ?
WHERE id = ?;
